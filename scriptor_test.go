package goredis_test

import (
	"testing"

	. "github.com/adimax2953/go-redis/src"

	goredis "github.com/adimax2953/go-redis"
	"github.com/stretchr/testify/assert"
)

var (
	scriptDefinition = "TGaming|0.0.0"
	dbKey            = "2"
	projectKey       = "minigame1"
	tagKey           = "game"
	keys             = []string{
		dbKey,
		projectKey,
		tagKey,
	}
	args = []string{
		"time",
		"一起寫",
	}
)

var (
	scripts = map[string]string{
		NewStringID: NewStringTemplate,
	}
)

func goredis_NewDB(host string, port int, scr map[string]string) (*goredis.Scriptor, error) {
	// opt := &goscriptor.Option{
	// 	Host:     host,
	// 	Port:     port,
	// 	Password: "",
	// 	DB:       0,
	// 	PoolSize: 1,
	// }

	opt := &goredis.Option{
		Host:     "192.168.10.183",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}

	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &scr)
	return scriptor, err
}

func goredis_New(host string, port int, scr map[string]string, assert *assert.Assertions) (*goredis.Scriptor, error) {
	// opt := &goredis.Option{
	// 	Host:     host,
	// 	Port:     port,
	// 	Password: "",
	// 	DB:       0,
	// 	PoolSize: 1,
	// }

	opt := &goredis.Option{
		Host:     "192.168.10.183",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}

	redis := opt.Create()

	assert.NotNil(redis)

	scriptor, err := goredis.New(redis, 1, scriptDefinition, &scr)
	return scriptor, err
}

func goredis_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {
	var res interface{}
	var err error

	res, err = scriptor.Exec("return 'Hello, World!'", keys, args)
	assert.Nil(err)
	assert.Equal("Hello, World!", res.(string), "they should be equal")

	res, err = scriptor.Exec("error return 'Hello, World!'", keys, args)
	assert.NotNil(err)

	res, err = scriptor.ExecSha(NewStringID, keys)
	assert.Nil(err)
	assert.Equal("Hello, World!", res.(string), "they should be equal")

	res, err = scriptor.ExecSha(NewStringID+" not found", keys, args)
	assert.NotNil(err)
	assert.Equal("script not found.", err.Error(), "they should be equal")
}

func goredis_TestCaseScriptNil(scriptor *goredis.Scriptor, assert *assert.Assertions) {
	var res interface{}
	var err error

	res, err = scriptor.Exec("return 'Hello, World!'", keys, args)
	assert.Nil(err)
	assert.Equal("Hello, World!", res.(string), "they should be equal")

	res, err = scriptor.Exec("error return 'Hello, World!'", keys, args)
	assert.NotNil(err)

	res, err = scriptor.ExecSha(NewStringID, keys, args)
	assert.NotNil(err)
	assert.Equal("script not found.", err.Error(), "they should be equal")
}

func Test_goredis_NewDB(t *testing.T) {
	/*var scriptor *goredis.Scriptor
	var err error

	assert := assert.New(t)

	// Mock Redis
	s := MockRedisServer()
	assert.NotNil(s)
	defer s.Close()

	// scripts does not exist
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, nil)
	assert.Nil(err)
	goredis_TestCaseScriptNil(scriptor, assert)

	// scripts is empty
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, map[string]string{})
	assert.Nil(err)
	goredis_TestCaseScriptNil(scriptor, assert)

	// register scripts
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, scripts)
	assert.Nil(err)
	// run test cases
	goredis_TestCase(scriptor, assert)

	// scripts does not exist, and reload redis cache scripts
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, nil)
	assert.Nil(err)
	// run test cases
	goredis_TestCase(scriptor, assert)

	// scripts is empty, and reload redis cache scripts
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, map[string]string{})
	assert.Nil(err)
	// run test cases
	goredis_TestCase(scriptor, assert)

	// flushAll redis
	ok, err := scriptor.Client.FlushAll(scriptor.CTX).Result()
	assert.Nil(err)
	assert.Equal("OK", ok, "they should be equal")

	// scripts does not exist
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, nil)
	assert.Nil(err)
	goredis_TestCaseScriptNil(scriptor, assert)

	// scripts is empty
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, map[string]string{})
	assert.Nil(err)
	goredis_TestCaseScriptNil(scriptor, assert)

	// can re-register scripts
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, scripts)
	assert.Nil(err)
	// run test cases
	goredis_TestCase(scriptor, assert)*/
}

func Test_goredis_New(t *testing.T) {
	// var scriptor *goredis.Scriptor
	// var err error

	// assert := assert.New(t)

	// // Mock Redis
	// s := MockRedisServer()
	// assert.NotNil(s)
	// defer s.Close()

	// // scripts does not exist
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, nil, assert)
	// assert.Nil(err)
	// goredis_TestCaseScriptNil(scriptor, assert)

	// // scripts is empty
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, map[string]string{}, assert)
	// assert.Nil(err)
	// goredis_TestCaseScriptNil(scriptor, assert)

	// // register scripts
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, scripts, assert)
	// assert.Nil(err)
	// // run test cases
	// goredis_TestCase(scriptor, assert)

	// // scripts does not exist, and reload redis cache scripts
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, nil, assert)
	// assert.Nil(err)
	// // run test cases
	// goredis_TestCase(scriptor, assert)

	// // scripts is empty, and reload redis cache scripts
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, map[string]string{}, assert)
	// assert.Nil(err)
	// // run test cases
	// goredis_TestCase(scriptor, assert)

	// // flushAll redis
	// ok, err := scriptor.Client.FlushAll(scriptor.CTX).Result()
	// assert.Nil(err)
	// assert.Equal("OK", ok, "they should be equal")

	// // scripts does not exist
	// scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, nil)
	// assert.Nil(err)
	// goredis_TestCaseScriptNil(scriptor, assert)

	// // scripts is empty
	// scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, map[string]string{})
	// assert.Nil(err)
	// goredis_TestCaseScriptNil(scriptor, assert)

	// // can re-register scripts
	// scriptor, err = goredis_New(s.Host(), s.Server().Addr().Port, scripts, assert)
	// assert.Nil(err)
	// // run test cases
	// goredis_TestCase(scriptor, assert)
}
