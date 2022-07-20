package src

// DelString function - keys, args[] string
func (s *MyScriptor) DelString(keys, args []string) {

	s.Scriptor.ExecSha(DelStringID, keys, args)
}

// DelString - 寫入一個字串
const (
	DelStringID       = "DelString"
	DelStringTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelString
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local sender                                        = "DelString.lua"
		if not DBKey or DBKey=="" then
			return  {err="invalid argument 'DBKey'", sender=sender}
		end
		if not ProjectKey or ProjectKey=="" then
			return  {err="invalid argument 'ProjectKey'", sender=sender}
		end
		if not TagKey or TagKey=="" then
			return  {err="invalid argument 'TagKey'", sender=sender}
		end
		if not k1 or k1=="" then
			return  {err="invalid argument 'k1'", sender=sender}
		end
		if DBKey and ProjectKey and TagKey and k1 then
			redis.call("select",DBKey)

			local r1= redis.call('del',ProjectKey..":"..TagKey..":"..k1)

			return { tostring(r1) }
		end
    `
)
