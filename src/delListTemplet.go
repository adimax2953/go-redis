package Src

// DelList function - keys, args[] string
func (s *MyScriptor) DelList(keys, args []string) {

	s.Scriptor.ExecSha(DelListID, keys, args)
}

// DelList - 減少數值
const (
	DelListID       = "DelList"
	DelListTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelList
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v1} {v2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v1                                            = tonumber(ARGV[2])    
		local v2                                            = ARGV[3]
		local sender                                        = "DelList.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v1 and v2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
			redis.call('lrem',MAIN_KEY,v1,v2)
		end
    `
)
