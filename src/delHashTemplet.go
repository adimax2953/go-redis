package Src

// DelHash function - keys, args[] string
func (s *MyScriptor) DelHash(keys, args []string) {

	s.Scriptor.ExecSha(DelHashID, keys, args)
}

// DelHash - 減少數值
const (
	DelHashID       = "DelHash"
	DelHashTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelHash
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {k2}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "DelHash.lua"
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('hdel',MAIN_KEY , k2)
		end
    `
)
