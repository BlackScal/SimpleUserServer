package redis

import (
	"encoding/json"
	"errors"
	"time"
	"userserver/internal/pkg/cache"

	"github.com/go-redis/redis"
)

var (
	tokenKeyPrefix      = "token_"
	tokenExpireDuration time.Duration
)

type Redis struct {
	options redis.Options
	conn    *redis.Client
}

func NewRedis(addr string) *Redis {
	redis := Redis{}
	redis.options.Addr = addr
	return &redis
}

func (r *Redis) SetPoolSize(poolSize int) {
	r.options.PoolSize = poolSize
}

func (r *Redis) SetExpireDuration(d time.Duration) {
	tokenExpireDuration = d
}

func (r *Redis) SetDialTimout(d time.Duration) {
	r.options.DialTimeout = d
}

func (r *Redis) SetWriteTimeout(d time.Duration) {
	r.options.WriteTimeout = d
}

func (r *Redis) SetReadTimeout(d time.Duration) {
	r.options.ReadTimeout = d
}

func (r *Redis) SetDataBase(db int) {
	r.options.DB = db
}

func (r *Redis) Init() error {
	conn := redis.NewClient(&r.options)
	if conn == nil {
		return errors.New("Redis Init failed to create client.")
	}

	if _, err := conn.Ping().Result(); err != nil {
		return errors.New("Failed to connect redis server")
	}
	r.conn = conn
	return nil
}

func (r *Redis) Destroy() {
	if r.conn == nil {
		return
	}
	r.conn.Close()
}

func (r *Redis) GetToken(tokeninfo *cache.TokenCacheInfo) error {
	key := tokenKeyPrefix + tokeninfo.Token
	val, err := r.conn.Get(key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), tokeninfo)
}

func (r *Redis) SetToken(tokeninfo *cache.TokenCacheInfo) error {
	key := tokenKeyPrefix + tokeninfo.Token
	val, err := json.Marshal(tokeninfo)
	if err != nil {
		return err
	}
	expired := time.Second * time.Duration(tokenExpireDuration)
	_, err = r.conn.Set(key, val, expired).Result()
	return err
}

func (r *Redis) DelToken(tokeninfo *cache.TokenCacheInfo) error {
	key := tokenKeyPrefix + tokeninfo.Token
	_, err := r.conn.Del(key).Result()
	return err

}
