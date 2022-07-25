package Src

// DelListAll function - keys, args[] string
func (s *MyScriptor) DelListAll(keys, args []string) {

	s.Scriptor.ExecSha(DelListAllID, keys, args)
}

// DelListAll - 減少數值
const (
	DelListAllID       = "DelListAll"
	DelListAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelListAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "DelListAll.lua"
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('del',MAIN_KEY)
		end
    `
)
