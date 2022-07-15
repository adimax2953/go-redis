package src

import (
	logtool "github.com/adimax2953/log-tool"
)

// DecValue function - keys, args[] string - return int64 , error
func (s *MyScriptor) DecValue(keys, args []string) (int64, error) {
	_, err := s.Scriptor.ExecSha(DecValueID, keys, args)
	if err != nil {
		logtool.LogError("DecValue ExecSha Error", err)
		return 0, err
	}
	result := &RedisResult{}

	return result.CountDown, nil
}

// DecValue - 減少數值
const (
	DecValueID       = "DecValue"
	DecValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DecValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = tonumber(ARGV[3])
		local sender                                        = "DecValue.lua"
		
		local result
		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
		
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = redis.call('hget',MAIN_KEY,k2)
			redis.call('hset',MAIN_KEY,k2,tmp-v1)
			result = redis.call('hget',MAIN_KEY,k2)
			return {result}
		end
    `
)
