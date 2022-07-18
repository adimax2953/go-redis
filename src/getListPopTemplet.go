package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetListPop function - keys, args[] string - return RedisResult , error
// model :  L lpush list前 // R rpush list後
func (s *MyScriptor) GetListPop(keys, args []string) (*RedisResult, error) {
	res, err := s.Scriptor.ExecSha(GetListPopID, keys, args)
	if err != nil {
		logtool.LogError("GetListPop ExecSha Error", err)
		return nil, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()

	return result, nil
}

// GetListPop - 減少數值
const (
	GetListPopID       = "GetListPop"
	GetListPopTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetListPop
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {model}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local model                                         = ARGV[2]
		local sender                                        = "GetListPop.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and model then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local m =""
			if model == "L" then 
				m = "lpop"
			elseif model == "R" then 
				m = "rpop"
			end
			local r1 = redis.call(m,MAIN_KEY)
		
			return { r1 }
		end
    `
)
