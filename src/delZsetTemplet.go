package Src

// DelZset function - keys, args[] string
func (s *MyScriptor) DelZset(keys, args []string) {

	s.Scriptor.ExecSha(DelValueID, keys, args)
}

// DelZset - 寫入一個數字
const (
	DelZsetID       = "DelZset"
	DelZsetTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelZset
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local v2                                            = ARGV[2]
		local sender                                        = "DelZset.lua"
		
		if DBKey and ProjectKey and TagKey and k1 and v2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('zrem',MAIN_KEY,v2)
		end
    `
)
