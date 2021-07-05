package redis

import (
	"github.com/garyburd/redigo/redis"
	"medilane-api/config"
)

type Cli struct {
	conn redis.Conn
}

var instanceRedisCli *Cli = nil

func Connect(conf *config.Config) (conn *Cli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(Cli)
		var err error

		instanceRedisCli.conn, err = redis.Dial("tcp", conf.REDIS.URL)

		if err != nil {
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", conf.REDIS.Password); err != nil {
			instanceRedisCli.conn.Close()
			panic(err)
		}
	}

	return instanceRedisCli
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
