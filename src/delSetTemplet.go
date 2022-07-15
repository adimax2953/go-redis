package src

import (
	logtool "github.com/adimax2953/log-tool"
)

// DelSet function - keys, args[] string - return string , error
func (s *MyScriptor) DelSet(keys, args []string) (string, error) {
	_, err := s.Scriptor.ExecSha(DelSetID, keys, args)
	if err != nil {
		logtool.LogError("DelSet ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}

	return result.Value, nil
}

// DelSet - 減少數值
const (
	DelSetID       = "DelSet"
	DelSetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelSet
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = ARGV[2]
		local sender                                        = "DelSet.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			
			redis.call('srem',MAIN_KEY ,v1)
		end
    `
)
