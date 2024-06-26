package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// IncValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) IncValue(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(IncValueID, keys, args)
	if err != nil {
		logtool.LogError("IncValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.ValueInt64, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("IncValue Value Error", err)
		return 0, err
	}

	return result.ValueInt64, nil
}

// IncValue - 增加數字
const (
	IncValueID       = "IncValue"
	IncValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   IncValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "IncValue.lua"
		if not DBKey or DBKey=="" then
			return  {err="invalid argument 'DBKey'", sender=sender}
		end
		if not ProjectKey or ProjectKey=="" then
			return  {err="invalid argument 'ProjectKey'", sender=sender}
		end
		if not TagKey or TagKey=="" then
			return  {err="invalid argument 'TagKey'", sender=sender}
		end
		if not k1 or k1=="" then
			return  {err="invalid argument 'k1'", sender=sender}
		end
		if not k2 or k2=="" then
			return  {err="invalid argument 'k2'", sender=sender}
		end
		if not v1 or v1=="" then
			return  {err="invalid argument 'v1'", sender=sender}
		end
		
		---@return number 
		local function getTime()
			return redis.call("TIME")[1]
		end

		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then

			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			redis.call("select",DBKey)
			local result ={}			
			result = redis.call('hincrby',MAIN_KEY,k2,v1)
			redis.call('hset',MAIN_KEY, 'lastUpdateTime', getTime())
			return {result}
		end
    `
)
