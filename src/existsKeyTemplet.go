package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// ExistsKEY function - keys, args[] string - return string , error
func (s *MyScriptor) ExistsKEY(keys, args []string) (bool, error) {
	res, err := s.Scriptor.ExecSha(ExistsKEYID, keys, args)
	if err != nil {
		logtool.LogError("ExistsKEY ExecSha Error", err)
		return false, err
	}
	result := &RedisResult{}
	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	result.Value = reader.ReadString()
	if result.Value != "1" {
		return false, nil
	}

	return true, nil
}

// ExistsKEY - 檢查key存在否
const (
	ExistsKEYID       = "ExistsKEY"
	ExistsKEYTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   ExistsKEY
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "ExpireKEY.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1= "0"
			r1 =redis.call('EXISTS',MAIN_KEY)
			return { r1 }
		end
    `
)
