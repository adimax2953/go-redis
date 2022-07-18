package goredis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"
)

func Test_goredis_Script(t *testing.T) {
	var scriptor *goredis.Scriptor
	var err error

	assert := assert.New(t)

	// Mock Redis
	s := MockRedisServer()
	assert.NotNil(s)
	defer s.Close()

	// scripts does not exist
	scriptor, err = goredis_NewDB(s.Host(), s.Server().Addr().Port, nil)
	assert.Nil(err)
	//script_TestCase(scriptor, assert)

	roomscript_TestCase(scriptor, assert)

}

func script_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     16379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "TGaming|0.0.1"
		dbKey            = "2"
		projectKey       = "minigame1"
		tagKey           = "game"
		keys             = []string{
			dbKey,
			projectKey,
			tagKey,
		}
		args = []string{
			"AAAaa",
			"一起寫181 ",
		}
	)

	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.Scripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.UpdateString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug(res)
}

func roomscript_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     16379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "TGaming|0.0.1"
		/*dbKey            = "2"
		projectKey       = "minigame1"
		tagKey           = "game"
		keys             = []string{
			dbKey,
			projectKey,
			tagKey,
		}
		args = []string{
			"AAAaa",
			"一起寫181 ",
		}*/
	)
	logtool.LogDebug("", src.Scripts)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.Scripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.RoomJoin("2", "test", "2311", "game1", "coin1", "test001", 20, 1, "220718", false, "")
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("", res)

	res2, err := myscript.RoomLeft("2", "test", "2311", "game1", "coin1", "test001")
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("", res2)
}
