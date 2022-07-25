package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// NewList function - keys, args[] string - return string , error
func (s *MyScriptor) NewList(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(NewListID, keys, args)
	if err != nil {
		logtool.LogError("NewList ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("NewList Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// NewList - 減少數值
const (
	NewListID       = "NewList"
	NewListTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   NewList
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {model} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local model                                         = ARGV[2]
		local v1                                            = ARGV[3]
		local sender                                        = "NewList.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and model and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local m =""
			if model == "L" then 
				m = "lpush"
			elseif model == "R" then 
				m = "rpush"
			end
			local r1 = ""
			local tmp = redis.call(m,MAIN_KEY ,v1)
			if tmp>= 1 then
				r1= v1
			end
			return { r1 }
		end
    `
)
