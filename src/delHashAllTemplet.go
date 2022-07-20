package src

// DelHashAll function - keys, args[] string
func (s *MyScriptor) DelHashAll(keys, args []string) {

	s.Scriptor.ExecSha(DelHashAllID, keys, args)
}

// DelHashAll - 減少數值
const (
	DelHashAllID       = "DelHashAll"
	DelHashAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelHashAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "DelHashAll.lua"
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('del',MAIN_KEY)
		end
    `
)
