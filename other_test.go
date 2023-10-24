package goredis_test

import (
	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	Src "github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"
)

func Inc_Base62_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"aaa",
		"test",
		"zzzzz",
	}
	res, err := myscript.IncBase62(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncBase62", res)

}

func Exist_Key_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
		"aaa",
		"test",
		"zzzzz",
	}
	res, err := myscript.ExistsKEY(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ExistsKEY", res)

}
func Flush_DB_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "11"
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
	res, err := myscript.FlushDB(keys)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("FlushDB", res)

}
func Scan_DB_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "r-gs5qc5rlcax1uyjynvpd.redis.singapore.rds.aliyuncs.com",
		Port:     6379,
		Password: "qxp_PEZ4cqw8ehr3wfa",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "0"
		Count            = "0"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.ScanKey([]string{dbKey, Count}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ScanKey", res)
}

func Key_Type_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "Taijc@888",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "0"
		Count            = "asdasd"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.KeyType([]string{dbKey, Count}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("KeyType", res)
}
func Hset_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "r-gs5qc5rlcax1uyjynvpd.redis.singapore.rds.aliyuncs.com",
		Port:     6379,
		Password: "qxp_PEZ4cqw8ehr3wfa",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "0"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}

	opt2 := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "Taijc@888",
		DB:       15,
		PoolSize: 3,
	}
	scriptor2, err := goredis.NewDB(opt2, opt2.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	myscript2 := &Src.MyScriptor{
		Scriptor: scriptor2,
	}

	res, err := myscript.HGetAll([]string{dbKey, "auto-ssl"}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	argsmap := make(map[string]interface{})
	for _, v := range *res {
		logtool.LogDebug(v.Key, v.Value)

		argsmap[v.Key] = v.Value
	}

	myscript2.HSet([]string{dbKey, "auto-ssl"}, argsmap)
}
