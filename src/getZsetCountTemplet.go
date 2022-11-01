package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// DelSet function - keys, args[] string
func (s *MyScriptor) GetZsetCount(keys, args []string) (int64, error) {

	res, err := s.Scriptor.ExecSha(GetZsetCountID, keys, args)
	if err != nil {
		logtool.LogError("GetZsetCount ExecSha Error", err)
		return 0, err
	}

	count, err := goredis.NewRedisReplyValue(res).AsInt64(0)
	if err != nil {
		logtool.LogError("GetZsetCount Value Error", err)
		return 0, err
	}
	return count, nil
}

// GetZsetAll - 寫入一個字串
const (
	GetZsetCountID       = "GetZsetCount"
	GetZsetCountTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetZsetCount
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = tonumber(ARGV[2])
		local k3                                            = tonumber(ARGV[3])
		local sender                                        = "GetZsetCount.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			local tmp = redis.call('zcount',MAIN_KEY , k2, k3)
		
			return tmp
		end
	`
)
