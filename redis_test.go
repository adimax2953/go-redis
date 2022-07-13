package goredis_test

import (
	"context"
	"goredis"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis/v2"
)

func MockRedisServer() *miniredis.Miniredis {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	s.FlushAll()

	return s
}

func Test_redis(t *testing.T) {
	var err error

	assert := assert.New(t)

	mockRedisServer := MockRedisServer()
	assert.NotNil(mockRedisServer)
	defer mockRedisServer.Close()

	opt := &goredis.Option{
		Host:     mockRedisServer.Host(),
		Port:     mockRedisServer.Server().Addr().Port,
		Password: "",
		DB:       0,
		PoolSize: 1,
	}

	redis := opt.Create()
	assert.NotNil(redis)

	_, err = redis.Ping(context.Background()).Result()
	assert.Nil(err)
}
