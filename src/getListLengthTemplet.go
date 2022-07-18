package src

import (
	goredis "github.com/adimax2953/go-redis"
	logtool "github.com/adimax2953/log-tool"
)

// GetListLength function - keys, args[] string - return int64 , error
func (s *MyScriptor) GetListLength(keys, args []string) (int64, error) {
	res, err := s.Scriptor.ExecSha(GetListLengthID, keys, args)
	if err != nil {
		logtool.LogError("GetListLength ExecSha Error", err)
		return 0, err
	}

	reader := goredis.NewRedisArrayReplyReader(res.([]interface{}))
	count, err := reader.ReadInt64(0)
	if err != nil {
		logtool.LogError("GetListLength Value Error", err)
		return 0, err
	}

	return count, nil
}

// GetListLength - 減少數值
const (
	GetListLengthID       = "GetListLength"
	GetListLengthTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   GetListLength
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "GetListLength.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			local r1 = "-1"
			local tmp = redis.call('llen',MAIN_KEY)
			if tmp~= nil and tmp~= false then
				r1= tmp
			else
				r1= ""
			end
			return { tostring(r1) }
		end
    `
)
