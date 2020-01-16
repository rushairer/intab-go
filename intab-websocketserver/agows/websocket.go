package agows

import (
	//"log"
	"encoding/json"
	"fmt"
	oauth2 "intab-core/oauth2"
	"log"
	"net/http"
	"reflect"
	"sync"

	"github.com/imdario/mergo"
	"github.com/rushairer/ago"
	"github.com/segmentio/ksuid"
	"gopkg.in/olahol/melody.v1"
)

const (
	//WebSocketCommandDisConnect 断开
	WebSocketCommandDisConnect = 0

	//WebSocketCommandConnect 连接
	WebSocketCommandConnect = 1

	//WebSocketCommandMsg 发送消息给自己
	WebSocketCommandMsg = 2

	//WebSocketCommandMsgChannel 发送消息给频道里的其他人
	WebSocketCommandMsgChannel = 3

	//WebSocketCommandMsgAll 发送消息给所有人
	WebSocketCommandMsgAll = 4
)

/*
* Tag:
* sys		系统通知
* conn      自己已经连接，客户端通过这条消息获取Message.ID
*
*
 */

//WebSocketMessage 消息类
type WebSocketMessage struct {
	ID   string `json:"id"`
	Type int    `json:"type"`
	Tag  string `json:"tag"`
	Data string `json:"data"`
}

//WebSocketServer WebSocker服务器类
type WebSocketServer struct {
	Messages map[*melody.Session]*WebSocketMessage
	Routes   Routes
}

//EncodeMsg 打包消息
func (wssrv *WebSocketServer) EncodeMsg(msg *WebSocketMessage) []byte {
	str, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	return str
}

//DecodeMsg 解包消息
func (wssrv *WebSocketServer) DecodeMsg(msgData []byte) WebSocketMessage {
	var wsMsg WebSocketMessage
	json.Unmarshal(msgData, &wsMsg)
	return wsMsg
}

func (wssrv *WebSocketServer) logMessageIDs() {
	/*
		for i, msg := range wssrv.Messages {
			log.Println("[", i, "]:", msg.ID)
		}
	*/
}

//AddRoute Add route
func (wssrv *WebSocketServer) AddRoute(controller ControllerInterface, pattern string, method int, methodName string) Routes {
	var routes Routes

	routes = map[string]*Route{
		pattern: &Route{
			Pattern:    pattern,
			Controller: controller,
			Method:     method,
			MethodName: methodName,
		},
	}

	if err := mergo.Merge(&wssrv.Routes, routes); err != nil {
		log.Println(err)
	}

	return routes
}

//Run 运行服务
func (wssrv *WebSocketServer) Run(addr string) {
	m := melody.New()
	// Maximum size in bytes of a message.
	m.Config.MaxMessageSize = 1024 * 1024
	// The max amount of messages that can be in a sessions buffer before it starts dropping them.
	m.Config.MessageBufferSize = 1024

	wssrv.Messages = make(map[*melody.Session]*WebSocketMessage)
	lock := new(sync.Mutex)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		oauth2.Server().SetAllowGetAccessRequest(true)
		tokenInfo, err := oauth2.Server().ValidationBearerToken(r)

		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(ago.Result401); err != nil {
				panic(err)
			}
		} else {
			//频道
			ch := r.FormValue("ch")
			m.HandleRequestWithKeys(w, r, map[string]interface{}{
				"uid":   tokenInfo.GetUserID(),
				"token": tokenInfo.GetAccess(),
				"ch":    ch,
			})
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		lock.Lock()
		for q, msg := range wssrv.Messages {
			if q.Keys["ch"] == s.Keys["ch"] {
				s.Write(wssrv.EncodeMsg(msg))
			}
		}
		newMsg := WebSocketMessage{
			ID:   ksuid.New().String(),
			Type: WebSocketCommandConnect,
			Tag:  "sys",
			Data: fmt.Sprintf("{\"uid\": %s}", s.Keys["uid"].(string)),
		}

		wssrv.Messages[s] = &newMsg
		log.Println("connected numbers:", len(wssrv.Messages))
		wssrv.logMessageIDs()
		//对通知自己的消息修改key为conn
		//复制变量
		selfMsg := newMsg
		selfMsg.Tag = "conn"
		s.Write(wssrv.EncodeMsg(&selfMsg))
		lock.Unlock()
	})

	m.HandleDisconnect(func(s *melody.Session) {
		lock.Lock()
		disMsg := wssrv.Messages[s]
		disMsg.Type = WebSocketCommandDisConnect
		disMsg.Tag = "sys"
		disMsg.Data = fmt.Sprintf("{\"uid\": %s}", s.Keys["uid"].(string))
		wssrv.Messages[s] = disMsg
		m.BroadcastOthers(wssrv.EncodeMsg(wssrv.Messages[s]), s)
		delete(wssrv.Messages, s)
		log.Println("connected numbers:", len(wssrv.Messages))
		wssrv.logMessageIDs()
		lock.Unlock()
	})

	m.HandleMessage(func(s *melody.Session, message []byte) {
		wsMsg := wssrv.DecodeMsg(message)
		msg := wssrv.Messages[s]

		lock.Lock()
		if len(wsMsg.ID) > 0 {
			if msg.ID == wsMsg.ID {
				//log.Println("This is a message from ", msg.ID)
				wssrv.Messages[s] = &wsMsg

				tag := wsMsg.Tag
				route, ok := wssrv.Routes[tag]

				if !ok {
					str, err := json.Marshal(ago.Result404)
					if err != nil {
						panic(err)
					}
					wsMsg.Data = string(str)
					m.Broadcast(wssrv.EncodeMsg(wssrv.Messages[s]))
				} else {

					t := reflect.Indirect(reflect.ValueOf(route.Controller)).Type()
					controller := reflect.New(t)

					//Call Controller Init Method
					initMethod := controller.MethodByName("Init")
					vars := make([]reflect.Value, 2)
					vars[0] = reflect.ValueOf(s)
					vars[1] = reflect.ValueOf(wsMsg)
					initMethod.Call(vars)

					//Call Controller Prepare Method
					vars = make([]reflect.Value, 0)
					method := controller.MethodByName("Prepare")
					method.Call(vars)

					vars = make([]reflect.Value, 1)
					vars[0] = reflect.ValueOf([]byte{})
					method = controller.MethodByName(route.MethodName)
					results := method.Call(vars)

					wsMsg.Type = route.Method

					switch wsMsg.Type {
					case WebSocketCommandMsg:
						m.BroadcastFilter(results[0].Interface().([]byte), func(q *melody.Session) bool {
							return q == s
						})
					case WebSocketCommandMsgChannel:
						m.BroadcastFilter(results[0].Interface().([]byte), func(q *melody.Session) bool {
							return q.Keys["ch"] == s.Keys["ch"]
						})
					case WebSocketCommandMsgAll:
						m.Broadcast(results[0].Interface().([]byte))
					}
				}

			} else {
				log.Println(msg.ID, " vs ", wsMsg.ID)
			}
		}
		lock.Unlock()
	})

	log.Println("Running server " + addr + " ...")
	http.ListenAndServe(addr, nil)
}
