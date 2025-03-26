package goredis

import (
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// Option - Redis Option
type Option struct {
	Host             string
	Port             int
	Password         string
	DB               int
	PoolSize         int
	ScriptDefinition string

	PoolFIFO        bool
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
	PoolTimeout     time.Duration
	MinIdleConns    int
	MaxActiveConns  int
	MaxLifetime     time.Duration
}

// Create - create a new redis descriptor
func (opt *Option) Create() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:            opt.Host + ":" + strconv.Itoa(opt.Port),
		Password:        opt.Password,
		DB:              opt.DB,
		PoolSize:        opt.PoolSize,
		ReadTimeout:     -1,
		PoolFIFO:        opt.PoolFIFO,
		MaxIdleConns:    opt.MaxIdleConns,
		ConnMaxIdleTime: opt.ConnMaxIdleTime,
		PoolTimeout:     opt.PoolTimeout,
		MinIdleConns:    opt.MinIdleConns,
		MaxActiveConns:  opt.MaxActiveConns,
		ConnMaxLifetime: opt.MaxLifetime,
	})
}
