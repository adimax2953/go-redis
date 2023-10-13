package goredis_test

import (
	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	Src "github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"
)

func value_inc_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value",
		"50",
	}
	res, err := myscript.IncValue(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ValueInc", res)
}

func value_dec_nag_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value",
		"100",
	}
	res, err := myscript.DecNegativeValue(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("DecNegativeValue", res)
}
func value_dec_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value",
		"100",
	}
	res, err := myscript.DecValue(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ValueDec", res)
}

func value_get_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value3",
	}
	res, err := myscript.GetValue(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ValueGet", res)
}
func value_get_all_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value",
	}
	res, err := myscript.GetValueAll(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ValueGetAll", *res)
}
func value_del_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"valuetest",
		"Value3",
	}
	myscript.DelValue(keys, args)

	logtool.LogDebug("ValueDel")
}

func value_inc_fixed_ttl_map_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
			"100",
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
		HttpAddr    int64 `redis:"1"`
		WsAddr      int64 `redis:"2"`
		WsPath      int64 `redis:"3"`
		SwaggerAddr int64 `redis:"4"`
	}

	args := Setting{
		HttpAddr:    123,
		WsAddr:      1,
		WsPath:      321,
		SwaggerAddr: 2,
	}

	res, err := myscript.IncValueBatchFixedTTL(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncValueBatchFixedTTL", *res)

	argsmap := make(map[string]interface{})
	argsmap["5"] = 1232133
	argsmap["6"] = 333
	argsmap["7"] = 222

	res, err = myscript.IncValueBatchFixedTTL(keys, argsmap)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncValueBatchFixedTTL", *res)
}
func value_inc_map_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		HttpAddr    int64 `redis:"1"`
		WsAddr      int64 `redis:"2"`
		WsPath      int64 `redis:"3"`
		SwaggerAddr int64 `redis:"4"`
	}

	args := Setting{
		HttpAddr:    123,
		WsAddr:      1,
		WsPath:      321,
		SwaggerAddr: 2,
	}

	res, err := myscript.IncValueBatch(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncValueBatch", *res)

	argsmap := make(map[string]interface{})
	argsmap["5"] = 1232133
	argsmap["6"] = 333
	argsmap["7"] = 222

	res, err = myscript.IncValueBatch(keys, argsmap)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncValueBatch", *res)
}

func value_dec_map_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		HttpAddr    int64 `redis:"1"`
		WsAddr      int64 `redis:"2"`
		WsPath      int64 `redis:"3"`
		SwaggerAddr int64 `redis:"4"`
	}

	args := Setting{
		HttpAddr:    123,
		WsAddr:      1,
		WsPath:      321,
		SwaggerAddr: 2,
	}

	res, err := myscript.DecValueBatch(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("DecValueBatch", *res)

	argsmap := make(map[string]interface{})
	argsmap["5"] = 1232133
	argsmap["6"] = 333
	argsmap["7"] = 222

	res, err = myscript.DecValueBatch(keys, argsmap)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("DecValueBatch", *res)
}
