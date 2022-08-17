package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// TakeValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) TakeValue(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(TakeValueID, keys, args)
	if err != nil {
		logtool.LogError("TakeValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("TakeValue Value Error", err)
		return 0, err
	}

	return result.ValueInt64, nil
}

// TakeValue - 取得數字
const (
	TakeValueID       = "TakeValue"
	TakeValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   TakeValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "TakeValue.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local result
			
			local POOL = redis.call('hget',MAIN_KEY,k2)
			local tmp_money = 0
		
			if POOL~=nil and POOL~="" and POOL~=false then
				tmp_money = POOL * v1 /100
			else
				POOL = 0
				redis.call('hset',MAIN_KEY, k2, 0 )  
			end
		
			redis.call('hset',MAIN_KEY, k2, POOL - tmp_money )  
			result = tostring(tmp_money)
			
			return {result}
		end
    `
)
