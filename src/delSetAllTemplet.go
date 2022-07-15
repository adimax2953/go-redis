package src

import (
	logtool "github.com/adimax2953/log-tool"
)

// DelSetAll function - keys, args[] string - return string , error
func (s *MyScriptor) DelSetAll(keys, args []string) (string, error) {
	_, err := s.Scriptor.ExecSha(DelSetAllID, keys, args)
	if err != nil {
		logtool.LogError("DelSetAll ExecSha Error", err)
		return "", err
	}
	result := &RedisResult{}

	return result.Value, nil
}

// DelSetAll - 減少數值
const (
	DelSetAllID       = "DelSetAll"
	DelSetAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelSetAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "DelSetAll.lua"
		
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('del',MAIN_KEY)
		end
    `
)
