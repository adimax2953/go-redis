package src

import (
	"runtime/debug"

	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// HGetAll function - keys, args[] string - return *[]HashResult , error
func (s *MyScriptor) HGetAll(keys, args []string) (*[]RedisResult, error) {
	defer func() {
		e := recover()
		if e != nil {
			debug.PrintStack()
			return
		}
	}()
	if s == nil {
		logtool.LogError("HGetAll empty s")
	}
	if s.Scriptor == nil {
		logtool.LogError("HGetAll empty scriptor")
	}
	res, err := s.Scriptor.ExecSha(HGetAllID, keys, args)
	if err != nil {
		logtool.LogError("HGetAll ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count := len(res.([]interface{})) / 2
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Key = reader.ReadString()
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("HGetAll Value Error", err)
		}
	}

	return &result, nil
}

// HGetAll - 減少數值
const (
	HGetAllID       = "HGetAll"
	HGetAllTemplate = `
	--[[
		Author      :   Adima.Tsai
		Description :   HGetAll
		EVALSHA  <script_sha1> 0 {DBKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local k1                                            = KEYS[2]
		local sender                                        = "HGetAll.lua"
		
		if DBKey and k1 then
		
			redis.call("select",DBKey)
			local r1= "-1"
			local Tmp = redis.call('hgetall',k1 )
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return r1
		end
    `
)
