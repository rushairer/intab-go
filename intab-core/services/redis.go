package services

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var _redisInstance *redis.Pool

//InitRedis 初始化Redis服务
func InitRedis(conf map[string]string) {
	if _redisInstance == nil {

		redisPool := &redis.Pool{
			MaxIdle:     10,
			MaxActive:   100,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				return dial("tcp", conf["host"]+":"+conf["port"], conf["password"])
			},
		}
		_redisInstance = redisPool
	}
}

//Redis 获得Redis服务实例单例
func Redis() redis.Conn {
	return _redisInstance.Get()
}

//CloseRedis 关闭数据库连接
func CloseRedis() {
	Redis().Close()
}

func dial(network, address, password string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, errAuth := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, errAuth
		}
	}
	return c, err
}
