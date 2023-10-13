package goredis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	goredis "github.com/adimax2953/go-redis"
	Src "github.com/adimax2953/go-redis/src"
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
	//Inc_Base62_TestCase(scriptor, assert)

	// room test
	//room_join_TestCase(scriptor, assert)
	//room_left_TestCase(scriptor, assert)
	//room_join_bot_TestCase(scriptor, assert)
	// room_left_bot_TestCase(scriptor, assert)
	//room_left_single_TestCase(scriptor, assert)
	// room_list_TestCase(scriptor, assert)
	// room_player_TestCase(scriptor, assert)
	//room_id_player_TestCase(scriptor, assert)

	//string test
	// string_new_TestCase(scriptor, assert)
	// string_update_TestCase(scriptor, assert)
	// string_get_TestCase(scriptor, assert)
	// string_del_TestCase(scriptor, assert)
	// string_update_ttl_TestCase(scriptor, assert)

	//Hash test
	// hash_new_TestCase(scriptor, assert)
	// hash_update_TestCase(scriptor, assert)
	// hash_get_TestCase(scriptor, assert)
	// hash_get_all_TestCase(scriptor, assert)
	// hash_del_TestCase(scriptor, assert)
	// hash_del_all_TestCase(scriptor, assert)
	// hash_update_list_TestCase(scriptor, assert)
	// hash_get_normal_TestCase(scriptor, assert)
	// hash_update_map_TestCase(scriptor, assert)

	//list test
	// list_new_TestCase(scriptor, assert)
	// list_update_TestCase(scriptor, assert)
	// list_get_TestCase(scriptor, assert)
	// list_get_length_TestCase(scriptor, assert)
	// list_get_all_TestCase(scriptor, assert)
	// list_get_pop_TestCase(scriptor, assert)
	// list_del_TestCase(scriptor, assert)
	// list_del_all_TestCase(scriptor, assert)

	//set test
	//set_new_TestCase(scriptor, assert)
	// set_update_TestCase(scriptor, assert)
	// set_get_TestCase(scriptor, assert)
	// set_get_random_TestCase(scriptor, assert)
	// set_get_all_TestCase(scriptor, assert)
	// set_del_TestCase(scriptor, assert)
	// set_del_all_TestCase(scriptor, assert)

	//zset test
	// zset_new_TestCase(scriptor, assert)
	// zset_update_TestCase(scriptor, assert)
	// zset_get_TestCase(scriptor, assert)
	// zset_get_all_TestCase(scriptor, assert)
	// zset_del_TestCase(scriptor, assert)
	// zset_del_all_TestCase(scriptor, assert)

	//value test
	// value_inc_TestCase(scriptor, assert)
	// value_dec_TestCase(scriptor, assert)
	// value_get_TestCase(scriptor, assert)
	// value_get_all_TestCase(scriptor, assert)
	// value_del_TestCase(scriptor, assert)
	//value_dec_nag_TestCase(scriptor, assert)
	//value_inc_map_TestCase(scriptor, assert)
	//value_dec_map_TestCase(scriptor, assert)
	value_inc_fixed_ttl_map_TestCase(scriptor, assert)

	//Exist_Key_TestCase(scriptor, assert)
	//Flush_DB_TestCase(scriptor, assert)

}

func script_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions) {

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
	res, err := myscript.UpdateString(keys, args)
	if err != nil {
		logtool.LogFatal(err.Error())
	}
	logtool.LogDebug(res)
}
