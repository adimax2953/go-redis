package src

// DelZsetAll function - keys, args[] string
func (s *MyScriptor) DelZsetAll(keys, args []string) {

	s.Scriptor.ExecSha(DelZsetAllID, keys, args)
}

// DelZsetAll - 寫入一個數字
const (
	DelZsetAllID       = "DelZsetAll"
	DelZsetAllTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelZsetAll
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "DelZsetAll.lua"
		if DBKey and ProjectKey and TagKey and k1 then
			local MAIN_KEY = ProjectKey..":"..TagKey..":"..k1
		
			redis.call("select",DBKey)
		
			redis.call('del',MAIN_KEY)
		end
    `
)
