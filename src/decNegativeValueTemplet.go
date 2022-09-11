package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// DecNegativeValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) DecNegativeValue(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(DecNegativeValueID, keys, args)
	if err != nil {
		logtool.LogError("DecValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("DecValue Value Error", err)
		return 0, err
	}

	return result.ValueInt64, nil
}

// DecNegativeValue - 減少數值
const (
	DecNegativeValueID       = "DecNegativeValue"
	DecNegativeValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DecNegativeValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "DecNegativeValue.lua"
		---@return number 
		local function getTime()
			return redis.call("TIME")[1]
		end

		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
		
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = redis.call('hget',MAIN_KEY,k2)
			local v2 =tonumber(tmp)
			if v2 == nil then
				v2=0
			end
			local result = {-1}
			redis.call('hset',MAIN_KEY,k2,v2 - v1)
			result = redis.call('hget',MAIN_KEY,k2)
			redis.call("hset",MAIN_KEY,"lastUpdateTime",getTime())
			

			return {result}
		end
    `
)
