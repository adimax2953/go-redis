package goredis_test


import (
	"testing"
	"github.com/stretchr/testify/assert"

	src "github.com/adimax2953/go-redis/src"
	logtool "github.com/adimax2953/log-tool"

	goredis "github.com/adimax2953/go-redis"
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
	script_TestCase(scriptor, assert)


}
func script_TestCase(scriptor *goredis.Scriptor, assert *assert.Assertions)  {

	opt := &goredis.Option{
		Host:     "192.168.56.1",
		Port:     16379,
		Password: "",
		DB:       1,
		PoolSize: 3,
	}
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
    
    scripts := map[string]string{
        src.UpdateStringID: src.UpdateStringTemplate,
        src.GetStringID: src.GetStringTemplate,
        src.NewStringID: src.NewStringTemplate,
        src.DelStringID: src.DelStringTemplate,
    }

    scriptor, err := goredis.NewDB(opt, 1, scriptDefinition, &scripts)
    if err != nil {
		logtool.LogFatal(err.Error())
    }	

    myscript := &src.MyScriptor{
        Scriptor: scriptor,
    }
    res, err := myscript.UpdateString(keys,args)
    if err != nil {
        logtool.LogFatal(err.Error())
    }
    logtool.LogDebug(res)
}