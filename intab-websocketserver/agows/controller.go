package agows

import (
	melody "gopkg.in/olahol/melody.v1"
)

//ControllerInterface   Base controller interface
type ControllerInterface interface {
	Init(s *melody.Session, msg WebSocketMessage)
}

//Controller Base controller
type Controller struct {
	Session *melody.Session
	Msg     WebSocketMessage
}

//Init Init controller
func (c *Controller) Init(s *melody.Session, msg WebSocketMessage) {
	c.Session = s
	c.Msg = msg
}

//Prepare Run this before all actions
func (c *Controller) Prepare() {

}
