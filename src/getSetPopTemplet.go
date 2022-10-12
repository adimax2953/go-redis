package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetSetPop function - keys, args[subkey,request datacount] string - return *[]RedisResult , error
func (s *MyScriptor) GetSetPop(keys, args []string) (*[]RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetSetPopID, keys, args)
	if err != nil {
		logtool.LogError("GetSetPop ExecSha Error", err)
		return nil, err
	}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))

	count := len(res.([]interface{}))
	result := make([]RedisResult, count)

	for i := 0; i < count; i++ {
		r := &RedisResult{}
		r.Value = reader.ReadString()
		result[i] = *r
		if err != nil {
			logtool.LogError("GetSetPop Value Error", err)
		}
	}
	return &result, nil
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
		local sender                                        = "GetSetPop.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local tmp = ""
			tmp = redis.call('SPOP',MAIN_KEY)
			
			return tmp 
		end
    `
)
