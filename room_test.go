package goredis_test

import (
	"strconv"

	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	"github.com/adimax2953/go-redis/Src"
	logtool "github.com/adimax2953/log-tool"
)

func room_list_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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

	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	res, err := myscript.RoomList(keys)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("RoomList:", res)
}

func room_player_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	res, err := myscript.RoomPlayer(keys, "2311", "game1", "coin1")
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("RoomPlayer", res)

}

func room_join_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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

	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	for i := 0; i < 10; i++ {
		res, err := myscript.RoomJoin(keys, "2311", "game1", "coin1", "test00"+strconv.Itoa(i), 20, 20, "220719", false, "")
		if err != nil {
			logtool.LogFatal(err.Error())
		}
		logtool.LogDebug("RoomJoin", res)
	}
}

func room_left_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	for i := 0; i < 5; i++ {
		res, err := myscript.RoomLeft(keys, "2311", "game1", "coin1", "test00"+strconv.Itoa(i))
		if err != nil {
			logtool.LogFatal(err.Error())
		}
		logtool.LogDebug("RoomLeft", res)
	}
}
func room_join_bot_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	for i := 10; i < 20; i++ {
		res, err := myscript.RoomJoin(keys, "2311", "game1", "coin1", "test00"+strconv.Itoa(i), 1, 20, "220719", true, "220719000000000"+strconv.Itoa(i-9))
		if err != nil {
			logtool.LogFatal(err.Error())
		}
		logtool.LogDebug("RoomJoinBot", res)
	}
}

func room_left_bot_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     6379,
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
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	for i := 10; i < 20; i++ {
		res, err := myscript.RoomLeft(keys, "2311", "game1", "coin1", "test00"+strconv.Itoa(i))
		if err != nil {
			logtool.LogFatal(err.Error())
		}
		logtool.LogDebug("RoomLeftBot", res)
	}
}
