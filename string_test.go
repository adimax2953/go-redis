package goredis_test

import (
	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"
)

func string_new_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"test",
		"NewString",
	}
	res, err := myscript.NewString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("NewString", res)

}

func string_update_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"test",
		"UpdateString",
	}
	res, err := myscript.UpdateString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateString", res)
}

func string_get_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"test",
		"GetString",
	}
	res, err := myscript.GetString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("GetString", res)
}
func string_del_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}

	args := []string{
		"test",
		"DelString",
	}
	myscript.DelString(keys, args)

	logtool.LogDebug("DelString")
}

func string_update_ttl_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &src.MyScriptor{
		Scriptor: scriptor,
	}

	args := []string{
		"test",
		"UpdateTTLString",
	}
	res, err := myscript.UpdateTTLString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateTTLString", res)
}
