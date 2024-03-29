package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// UpdateHashList function - keys, args[] string - return string , error
func (s *MyScriptor) UpdateHashList(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(UpdateHashListID, keys, args)
	if err != nil {
		logtool.LogError("UpdateHashList ExecSha Error", err)
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
			logtool.LogError("UpdateHashList Value Error", err)
		}
	}

	return &result, nil
}

// UpdateHashList - 減少數值
const (
	UpdateHashListID       = "UpdateHashList"
	UpdateHashListTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   UpdateHashList
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local v1                                            = ARGV[3]
		local sender                                        = "UpdateHashList.lua"
		
		---@param str string
		---@param separator string
		---@return string[]
		local function split(str, separator)
			local t = {}
			for s in string.gmatch(str, "([^" .. separator .. "]+)") do
				table.insert(t, s)
			end
			return t
		end

		if DBKey and ProjectKey and TagKey and k1 and k2 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
			local keyparts = split(k2, "~")
			local valueparts = split(v1, "~")
			
		
			redis.call("select",DBKey)
			local count = #valueparts
			for i = 1 , count do
				redis.call('hset',MAIN_KEY , keyparts[i] ,valueparts[i])
			end

			local r1 = ""
			local Tmp = redis.call('hgetall',MAIN_KEY )
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return r1 
		end
    `
)
