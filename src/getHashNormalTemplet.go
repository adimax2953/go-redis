package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetHashNormal function - keys, args[] string - return string , error
func (s *MyScriptor) GetHashNormal(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetHashNormalID, keys, args)
	if err != nil {
		logtool.LogError("GetHashNormal ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("GetHashNormal Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// GetHashNormal - 減少數值
const (
	GetHashNormalID       = "GetHashNormal"
	GetHashNormalTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetHashNormal
		EVALSHA  <script_sha1> 0 {DBKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetHashNormal.lua"
		if DBKey and k1 and k2 then
			local MAIN_KEY = k1
		
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
