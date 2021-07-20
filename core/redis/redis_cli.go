package redisCon

import (
	"github.com/garyburd/redigo/redis"
	"medilane-api/config"
)

type Cli struct {
	conn redis.Conn
}

func Connect(conf *config.Config) *Cli {
	var err error

	conn, err := redis.Dial("tcp", conf.REDIS.URL)

	if err != nil {
		panic(err)
	}

	if _, err := conn.Do("AUTH", conf.REDIS.Password); err != nil {
		_ = conn.Close()
		panic(err)
	}

	return &Cli{
		conn: conn,
	}
}

func (redisCli *Cli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *Cli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}

func (redisCli *Cli) Close() error {
	return redisCli.conn.Close()
}
