package redisCon

import (
	"context"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

// for queue, cache (i.e: redis)
type MsgQueue interface {
	Init(addr string, pwd string, db int)
	Close()

	Ping(ctx context.Context) (string, error)
	LLen(ctx context.Context, key string) (int64, error)
	TTL(ctx context.Context, key string) (timeLeft time.Duration, err error)
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)

	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) (string, error)

	LPush(ctx context.Context, queue string, values ...interface{}) (int64, error)
	RPop(ctx context.Context, queue string) (string, error)
	LPop(ctx context.Context, queue string) (string, error)
	LRange(ctx context.Context, queueName string, start, stop int64) ([]string, error)
	Clear(context.Context, string) (int64, error)
}

var once sync.Once

type RedisConnector struct {
	Client *redis.Client
}

// singleton for api
var singletonRedisConnector MsgQueue

func GetInstance() MsgQueue {
	once.Do(func() { // <-- atomic, does not allow repeating
		singletonRedisConnector = &RedisConnector{}
	})
	return singletonRedisConnector
}

func SetInstance(obj MsgQueue) {
	singletonRedisConnector = obj
}

//NewClient new client for worker
func NewClient(redisURL string, redisPassword string, redisDB int) MsgQueue {
	rdClient := &RedisConnector{}
	rdClient.Init(redisURL, redisPassword, redisDB)
	return rdClient
}

func (obj *RedisConnector) Init(redisHost string, password string, db int) {
	obj.Client = redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
}

func (obj *RedisConnector) Close() {
	err := obj.Client.Close()
	if err != nil {
		log.Fatal(err)
	}
	obj.Client = nil
}

// implement action
func (obj *RedisConnector) LPush(ctx context.Context, queueName string, values ...interface{}) (int64, error) {
	return obj.Client.LPush(ctx, queueName, values...).Result()
}

func (obj *RedisConnector) RPop(ctx context.Context, queueName string) (string, error) {
	return obj.Client.RPop(ctx, queueName).Result()
}

func (obj *RedisConnector) LPop(ctx context.Context, queueName string) (string, error) {
	return obj.Client.LPop(ctx, queueName).Result()
}

func (obj *RedisConnector) LRange(ctx context.Context, queueName string, start, stop int64) ([]string, error) {
	return obj.Client.LRange(ctx, queueName, start, stop).Result()
}

func (obj *RedisConnector) LLen(ctx context.Context, key string) (int64, error) {
	return obj.Client.LLen(ctx, key).Result()
}

func (obj *RedisConnector) Get(ctx context.Context, key string) (string, error) {
	return obj.Client.Get(ctx, key).Result()
}

func (obj *RedisConnector) Set(ctx context.Context, key string, values interface{}, ttl time.Duration) (string, error) {
	return obj.Client.Set(ctx, key, values, ttl).Result()
}

func (obj *RedisConnector) Ping(ctx context.Context) (string, error) {
	return obj.Client.Ping(ctx).Result()
}

func (obj *RedisConnector) TTL(ctx context.Context, key string) (timeLeft time.Duration, err error) {
	return obj.Client.TTL(ctx, key).Result()
}

func (obj *RedisConnector) Keys(ctx context.Context, pattern string) ([]string, error) {
	return obj.Client.Keys(ctx, pattern).Result()
}

func (obj *RedisConnector) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return obj.Client.Expire(ctx, key, expiration).Result()
}

func (obj *RedisConnector) Clear(ctx context.Context, key string) (int64, error) {
	return obj.Client.Del(ctx, key).Result()
}

//func (redisCli *Cli) SetValue(key string, value string, expiration ...interface{}) error {
//	_, err := redisCli.conn.Do("SET", key, value)
//
//	if err == nil && expiration != nil {
//		redisCli.conn.Do("EXPIRE", key, expiration[0])
//	}
//
//	return err
//}
//
//func (redisCli *Cli) GetValue(key string) (interface{}, error) {
//	return redisCli.conn.Do("GET", key)
//}
//
//func (redisCli *Cli) Close() error {
//	return redisCli.conn.Close()
//}
