package goredis_test

import (
	"fmt"

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

	tid := fmt.Sprintf("%s-%s-%08d", "231122", "40A", 0)
	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"TRA",
		"TID",
		tid,
	}
	res, err := myscript.IncBase62(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncBase62", res)

}

func Inc_Base10_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "15"
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

	tid := fmt.Sprintf("%010d", 1839999999)
	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	args := []string{
		"TRA",
		"TID",
		tid,
	}
	res, err := myscript.IncBase10(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("IncBase10", res)

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

func Expire_Key_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "12"
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
		"-1",
	}
	res, err := myscript.ExpireKEY(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ExpireKEY", res)

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
		Host:     "",
		Port:     6379,
		Password: "",
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
		Password: "",
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

func zadd_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "6"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	a, b, _ := myscript.ZAdd([]string{dbKey, "myzset"}, []string{"acac", "1"})
	logtool.LogDebug("Zadd", a, b)

	c, d, _ := myscript.ZAdd([]string{dbKey, "myzset"}, []string{"acdc", "1"})
	logtool.LogDebug("Zadd", c, d)

}
func Zrange_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "6"
		Key              = "myzset"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.ZRange([]string{dbKey, Key}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("Zrange", res)
}
func Hset_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "",
		Port:     6379,
		Password: "",
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
		Password: "",
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
func Scan_DB_Match_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "6"
		Key              = "DSG:bftest:CNY:302:1:55"
		Count            = "10000"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.ScanMatchKey([]string{dbKey, Key, Count}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ScanMatchKey", *res)
}

func Scan_DB_Matchs_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "2"
		Key              = "1_1_15"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.ScanMatchKeys([]string{dbKey, Key}, []string{})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("ScanMatchKeys", *res)
}

func Get_System_RTP_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

	opt := &goredis.Option{
		Host:     "103.103.81.12",
		Port:     6379,
		Password: "",
		DB:       15,
		PoolSize: 3,
	}
	var (
		scriptDefinition = "Bft|0.0.1"
		dbKey            = "2"
		//Key              = "1_1_15_401_2000"
	)
	scriptor, err := goredis.NewDB(opt, opt.DB, scriptDefinition, &Src.LuaScripts)
	if err != nil {
		logtool.LogFatal(err.Error())
	}

	myscript := &Src.MyScriptor{
		Scriptor: scriptor,
	}
	res, err := myscript.GetSystemRTP([]string{dbKey}, []string{"1_1_15_401_2000:Lifetime:System", "1_1_15_401_2000:202312_1:System", "1_1_15_401_2000:20231214_1:System"})
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug("GetSystemRTP", res)
}
