package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSetPop function - keys, args[subkey,request datacount] string - return *[]RedisResult , error
func (s *MyScriptor) GetSetPop(keys, args []string) (string, error) {
	res, err := s.Scriptor.ExecSha(GetSetPopID, keys, args)
	if err != nil {
		logtool.LogError("GetSetPop ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if err != nil {
		logtool.LogError("GetSet Value Error", err)
		return "", err
	}

	return result.Value, nil
}

// GetSetPop - 減少數值
const (
	GetSetPopID       = "GetSetPop"
	GetSetPopTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetSetPop
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "GetSetPop.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = ""
			tmp = redis.call('SPOP', MAIN_KEY , k2)
			
			return tmp 
		end
    `
)
