package go-redis

import (
	"strconv"

	"github.com/go-redis/redis/v9"
)

// Option - Redis Option
type Option struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

// Create - create a new redis descriptor
func (opt *Option) Create() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     opt.Host + ":" + strconv.Itoa(opt.Port),
		Password: opt.Password,
		DB:       opt.DB,
		PoolSize: opt.PoolSize,
	})
}
