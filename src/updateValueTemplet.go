package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) UpdateValue(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(UpdateValueID, keys, args)
	if err != nil {
		logtool.LogError("UpdateValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.CountDown, err = reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("UpdateValue Value Error", err)
		return 0, err
	}

	return result.CountDown, nil
}

// UpdateValue - 寫入一個數字
const (
	UpdateValueID       = "UpdateValue"
	UpdateValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "UpdateValue.lua"
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
		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
			redis.call("select",DBKey)
			local result ={}
			redis.call('hset',ProjectKey..":"..TagKey..":"..k1,k2,v1)
			result = redis.call('hget',ProjectKey..":"..TagKey..":"..k1,k2)
			return {result}
		end
    `
)
