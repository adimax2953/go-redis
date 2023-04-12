package goredis_test

import (
	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	Src "github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"
)

func hash_new_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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
	args := []string{
		"hashtest",
		"NewHashkey",
		"NewHash",
	}
	res, err := myscript.NewHash(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("NewHash", res)

}

func hash_update_map_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6378,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "14"
		projectKey       = "minigame1"
		tagKey           = "game"
		keys             = []string{
			dbKey,
			projectKey,
			tagKey,
			"mainkey",
		}
	)
	scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	type Setting struct {
		HttpAddr    string `redis:"httpaddr"`
		WsAddr      string `redis:"wsaddr"`
		WsPath      string `redis:"wspath"`
		SwaggerAddr string `redis:"swaggeraddr"`
	}

	args := Setting{
		HttpAddr:    "8.133.56.1",
		WsAddr:      "8.133.56.2",
		WsPath:      "/game",
		SwaggerAddr: "8.133.56.3",
	}

	res, err := myscript.UpdateHashBatch(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateHashBatch", *res)

	argsmap := make(map[string]interface{})
	argsmap["test1"] = "123213"
	argsmap["test2"] = "gooooooo13"
	argsmap["test3"] = "阿哩"

	res, err = myscript.UpdateHashBatch(keys, argsmap)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateHashBatch", *res)

}
func hash_update_list_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "14"
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
	args := []string{
		"hashtest",
		"playerid~score~cointype",
		"139755~1500~CNC",
	}
	res, err := myscript.UpdateHashList(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateHashList", *res)
}
func hash_update_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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
	args := []string{
		"hashtest",
		"UpdateHashkey",
		"UpdateHash",
	}
	res, err := myscript.UpdateHash(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("UpdateHash", res)
}

func hash_get_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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
	args := []string{
		"hashtest",
		"UpdateHash",
	}
	res, err := myscript.GetHash(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("GetHash", res)
}
func hash_get_all_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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
	args := []string{
		"hashtest",
		"GetHashAll",
	}
	res, err := myscript.GetHashAll(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("GetHashAll", *res)
}
func hash_del_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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

	args := []string{
		"hashtest",
		"UpdateHash",
	}
	myscript.DelHash(keys, args)

	logtool.LogDebug("DelHash")
}

func hash_del_all_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
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

	args := []string{
		"hashtest",
		"DelHashAll",
	}
	myscript.DelHashAll(keys, args)

	logtool.LogDebug("DelHashAll")
}

func hash_get_normal_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "192.168.10.182",
		Port:     6379,
		Password: "Taijc@888",
		DB:       11,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "4"
		keys             = []string{
			dbKey,
		}
	)
	scriptor, err := goredis.NewDB(opt, 11, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"TG:playerToRoom",
		"test1",
	}
	res, err := myscript.GetHashNormal(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("GetHashNormal", res)
}
