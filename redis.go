package gotuil

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Connection Redis连接
type Connection struct {
	Conn *redis.Pool
}

// InitRedis 初始化数据库连接
func InitRedis(setting *RedisSetting) *Connection {
	if setting == nil {
		return nil
	}
	port := 6379
	if setting.Port != 0 {
		port = setting.Port
	}
	conn := &redis.Pool{
		MaxIdle:     setting.MaxIdle,
		MaxActive:   setting.MaxActive,
		IdleTimeout: setting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", setting.Host, port))
			if err != nil {
				return nil, err
			}
			if setting.Password != "" {
				if _, err := c.Do("AUTH", setting.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &Connection{conn}
}

// SetString 写入字符串
func (c *Connection) SetString(key, val string, exp int) error {
	con := c.Conn.Get()
	defer con.Close()
	_, err := con.Do("SET", key, val)
	if exp != 0 {
		_, err = con.Do("EXPIRE", key, exp)
	}
	return err
}

// SetInterface 写入interface{}
func (c *Connection) SetInterface(key string, data interface{}, exp int) error {
	con := c.Conn.Get()
	defer con.Close()
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = con.Do("SET", key, val)
	if exp != 0 {
		_, err = con.Do("EXPIRE", key, exp)
	}
	return err
}

// GetString 获取字符串
func (c *Connection) GetString(key string) (string, error) {
	con := c.Conn.Get()
	defer con.Close()
	val, err := redis.String(con.Do("GET", key))
	return val, err
}

// GetInterface 获取interface{}
func (c *Connection) GetInterface(key string) (interface{}, error) {
	con := c.Conn.Get()
	defer con.Close()
	val, err := con.Do("GET", key)
	return val, err
}

// Exists 判断是否存在
func (c *Connection) Exists(key string) bool {
	con := c.Conn.Get()
	defer con.Close()
	exists, err := redis.Bool(con.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Delete 删除
func (c *Connection) Delete(key string) (bool, error) {
	con := c.Conn.Get()
	defer con.Close()
	return redis.Bool(con.Do("DEL", key))
}

// LikeDeletes like删除
func (c *Connection) LikeDeletes(key string) error {
	con := c.Conn.Get()
	defer con.Close()
	keys, err := redis.Strings(con.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = c.Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
