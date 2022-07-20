package src

// DelValue function - keys, args[] string
func (s *MyScriptor) DelValue(keys, args []string) {

	s.Scriptor.ExecSha(DelValueID, keys, args)
}

// DelValue - 寫入一個數字
const (
	DelValueID       = "DelValue"
	DelValueTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   DelValue
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1}
		--]]
		local DBKey                                         = tonumber(KEYS[1])
		local ProjectKey                                    = KEYS[2]
		local TagKey                                        = KEYS[3]
		local k1                                            = ARGV[1]
		local k2                                            = ARGV[2]
		local sender                                        = "DelValue.lua"
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
		if not k2 or k2=="" then
			return  {err="invalid argument 'k2'", sender=sender}
		end
		if DBKey and ProjectKey and TagKey and k1 and k2 then
			redis.call("select",DBKey)
			local result ={}
			redis.call('hdel',ProjectKey..":"..TagKey..":"..k1,k2)
			result = redis.call('hget',ProjectKey..":"..TagKey..":"..k1,k2)
			return {result}
		end
    `
)
