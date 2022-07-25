package Src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetHash function - keys, args[] string - return string , error
func (s *MyScriptor) GetHash(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetHashID, keys, args)
	if err != nil {
		logtool.LogError("GetHash ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("GetHash Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// GetHash - 減少數值
const (
	GetHashID       = "GetHash"
	GetHashTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetHash
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetHash.lua"
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= ""
			local Tmp = redis.call('hget',MAIN_KEY , k2)
			if Tmp~=nil and Tmp~="" and Tmp~=false then
				r1 = Tmp
			end
			return { r1 }
		end
    `
)
