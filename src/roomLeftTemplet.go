package src

import (
	logtool "github.com/adimax2953/log-tool"
	"github.com/pkg/errors"
)

// RoomLeft function - keys, args[] string - return string , error
func (s *MyScriptor) RoomLeft(
	redisDB string,
	projectID string,
	platformID string,
	gameID string,
	countryCode string,
	playerID string,

) (interface{}, error) {
	args := []string{
		redisDB,
		projectID,
		platformID,
		gameID,
		countryCode,
		playerID,
	}
	res, err := s.Scriptor.ExecSha(RoomLeftID, nil, args)
	if err != nil {
		logtool.LogError("RoomLeft ExecSha Error", err)
		return "", errors.WithStack(err)
	}

	return res, nil
}

// RoomLeft - 減少數值
const (
	RoomLeftID       = "RoomLeft"
	RoomLeftTemplate = `
	local redisDb = ARGV[1]
	local projectId = ARGV[2]
	local platformId = ARGV[3]
	local gameId = ARGV[4]
	local currency = ARGV[5]
	local playerId = ARGV[6]
	
	local scope = table.concat({projectId, platformId, gameId, currency}, "/")
	
	local function log(v)
		local s = ""
		if type(v) == 'table' then
			s = cjson.encode(v)
		else
			s = tostring(v)
		end
		redis.call("RPUSH", "log", s)
	end
	
	---@param list table
	---@return string
	local function joinBySlash(list)
		return table.concat(list, "/")
	end
	
	---@param list table
	---@return string
	local function makeKey(list)
		return joinBySlash({scope, unpack(list)})
	end
	
	---@return number 
	local function getTime()
		return redis.call("TIME")[1]
	end
	
	---@param key string
	---@return table
	local function hgetall(key)
		local entries = redis.call('HGETALL', key)
		local result = {}
		for i = 1, #entries, 2 do
			result[entries[i]] = entries[i + 1]
		end
		return result
	end
	
	---@param roomId string
	---@return table
	local function getPlayersInRoom(roomId)
		local players = hgetall(makeKey({"room", roomId, "playerToSeat"}))
		local list = {}
		for k, v in pairs(players) do
			table.insert(list, {userid = k, seatid = v, roomid = roomId})
		end
		if #list == 0 then
			return cjson.null
		end
		return list
	end
	
	---@param roomId string
	local function deleteRoom(roomId)
		redis.call(
			"PUBLISH", "RoomClose",
				table.concat(
					{platformId, gameId, currency, roomId, "RoomClose"}, "~"
				)
		)
		redis.call(
			"SREM", projectId .. "/rooms",
				table.concat({platformId, gameId, currency, roomId}, ":")
		)
	
		redis.call("ZREM", makeKey({"rooms"}), roomId)
		redis.call("ZREM", makeKey({"roomsAvailable"}), roomId)
		redis.call("DEL", makeKey({"room", roomId, "seatsAvailable"}))
		redis.call("DEL", makeKey({"room", roomId, "playerToSeat"}))
		redis.call("DEL", makeKey({"room", roomId}))
	end
	
	redis.call("SELECT", redisDb)
	
	local roomId = redis.call("HGET", makeKey({"playerToRoom"}), playerId)
	
	if not roomId then
		return "PLAYER_NOT_IN_A_ROOM"
	end
	
	local seatId = redis.call(
		"HGET", makeKey({"room", roomId, "playerToSeat"}), playerId
	)
	
	if not seatId then
		error("internal error: cannot find seatId.")
	end
	
	redis.call("SADD", makeKey({"room", roomId, "seatsAvailable"}), seatId)
	redis.call("HDEL", makeKey({"room", roomId, "playerToSeat"}), playerId)
	
	local currentPlayerCount = redis.call(
		"HINCRBY", makeKey({"room", roomId}), "currentPlayerCount", -1
	)
	
	local currentBotCount = redis.call(
		"HGET", makeKey({"room", roomId}), "currentBotCount"
	)
	
	if currentPlayerCount == 0 or (currentPlayerCount - currentBotCount == 0) then
		deleteRoom(roomId)
	else
		redis.call("ZADD", makeKey({"roomsAvailable"}), getTime(), roomId)
	end
	
	local players = getPlayersInRoom(roomId)
	
	redis.call(
		"PUBLISH", "RoomUpdate", cjson.encode(
			{
				projectId = projectId,
				platformId = platformId,
				gameId = gameId,
				currency = currency,
				roomId = roomId,
				players = players
			}
		)
	)
	
	redis.call("HDEL", makeKey({"playerToRoom"}), playerId)
	
	return "OK"	
    `
)
