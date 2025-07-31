package src

import (
	"strconv"

	"github.com/pkg/errors"
)

// RoomJoin function - args[] string - return string , error
func (s *MyScriptor) RoomCreateOrJoin(
	keys []string,
	platformID string,
	gameID string,
	countryCode string,
	playerID string,
	maxPlayerCount int,
	seatsCount int,
	date string, // YYMMDD
	isBot bool,
	roomId string,
) (string, error) {
	isBotString := ""
	if isBot {
		isBotString = "isBot"
	}
	args := []string{
		platformID,
		gameID,
		countryCode,
		playerID,
		strconv.Itoa(maxPlayerCount),
		strconv.Itoa(seatsCount),
		date,
		isBotString,
		roomId,
	}
	res, err := s.Scriptor.ExecSha(RoomCreateOrJoinID, keys, args)
	if err != nil {
		return "", errors.WithStack(err)
	}

	result := res.(string)

	return result, nil
}

// RoomCreateOrJoin - 減少數值
const (
	RoomCreateOrJoinID       = "RoomCreateOrJoin"
	RoomCreateOrJoinTemplate = `
	--[[
		Author      :   Adimax.Tsai
		Description :   RoomCreateOrJoin
		EVALSHA  <script_sha1> 0 {DBKey} {ProjectKey} {TagKey} {k1} {v2}
	--]]
	-- 配房功能
	local DBKey                                         = tonumber(KEYS[1])
	local ProjectKey                                    = KEYS[2]
	local TagKey                                        = KEYS[3]
	local platformId = ARGV[1]
	local gameId = ARGV[2]
	local countrycode = ARGV[3]
	local playerId = ARGV[4]
	local maxPlayerCount = tonumber(ARGV[5])
	local seatsCount = tonumber(ARGV[6])
	local date = ARGV[7] -- YYMMDD, this is used to generate roomId
	local isBot = ARGV[8] == "isBot"
	local roomId = ARGV[9]
	local scope = table.concat({ProjectKey, platformId, gameId, countrycode}, "/")
	
	if maxPlayerCount == nil or not (maxPlayerCount > 0) then
		error("maxPlayerCount should be a positive integer")
	end
	
	if seatsCount == nil or not (seatsCount > 0) then
		error("seatsCount should be a positive integer")
	end
		
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
	
	---@param key string
	---@param t table
	local function hset(key, t)
		local entries = {}
		for k, v in pairs(t) do
			table.insert(entries, k)
			table.insert(entries, v)
		end
		return redis.call('HSET', key, unpack(entries))
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
	
	---@return number 
	local function getTime()
		return redis.call("TIME")[1]
	end
		
	---@param roomId string
	---@return table
	local function getRoom(roomId)
		return hgetall(makeKey({"room", roomId}))
	end
			
	---@param playerId string
	---@param room table
	---@return table, string
	local function addPlayerToRoom(playerId, room, playerType)
		room.currentPlayerCount = room.currentPlayerCount + 1
		if isBot then
			room.currentBotCount = room.currentBotCount + 1
		end
		hset(makeKey({"room", room.roomId}), room)
		if room.currentPlayerCount == tonumber(room.maxPlayerCount) then
			redis.call("ZREM", makeKey({"roomsAvailable"}), room.roomId)
		end
	
		local seatId = redis.call(
			"SPOP", joinBySlash({scope, "room", room.roomId, "seatsAvailable"})
		)
	
		redis.call(
			"HSET", makeKey({"room", room.roomId, "playerToSeat"}), playerId, seatId
		)
	
		redis.call("HSET", makeKey({"playerToRoom"}), playerId, table.concat({playerType, seatId, room.roomId,ProjectKey}, ":"))
		redis.call("HSET",  platformId..":"..gameId..":".."playerToRoom", playerId,table.concat({playerType, seatId, room.roomId,ProjectKey}, ":"))

		return room, seatId
	end
	
	---@param roomId string
	---@param playerId string
	---@param seatId string
	---@param players table
	local function publishPlayerJoinRoom(roomId, playerId, seatId, players)
		redis.call(
			"PUBLISH", "RoomUpdate", cjson.encode(
				{
					ProjectKey= ProjectKey,
					platformId = platformId,
					gameId = gameId,
					countrycode = countrycode,
					roomId = roomId,
					players = players
				}
			)
		)
	end
	
	---@param playerId string
	---@return string
	local function getRoomIdOfPlayer(playerId)
		return redis.call("HGET", makeKey({"playerToRoom"}), playerId)
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
	
	---@return table
	local function createRoomById(id)
		local roomId = id
		local room = {
			roomId = roomId,
			currentPlayerCount = 0,
			currentBotCount = 0,
			maxPlayerCount = maxPlayerCount
		}
	
		redis.call(
			"SADD", ProjectKey.. "/rooms",
				table.concat({platformId, gameId, countrycode, roomId}, "/")
		)
		redis.call("ZADD", makeKey({"rooms"}), getTime(), roomId)
		hset(makeKey({"room", roomId}), room)
		redis.call("ZADD", makeKey({"roomsAvailable"}), getTime(), roomId)
	
		local seats = {}
		for i = 1, seatsCount do
			table.insert(seats, tostring(i))
		end
	
		redis.call(
			"SADD", makeKey({"room", room.roomId, "seatsAvailable"}), unpack(seats)
		)
	
		redis.call(
			"PUBLISH", "RoomOpen",
				table.concat({platformId, gameId, countrycode, roomId, "RoomOpen"}, "~")
		)
		return room
	end

	redis.call("select", DBKey)
	
	if not isBot and getRoomIdOfPlayer(playerId) then
		return cjson.encode(
			{
				error = {
					code = "PLAYER_ALREADY_IN_ROOM",
					message = "Player is already in a room."
				}
			}
		)
	end
	
	local room
	local playerType = "P"

	room = getRoom(roomId)

	if next(room) == nil then
		room = createRoomById(roomId)
		log(room)
	end
		
	local room, seatId = addPlayerToRoom(playerId, room, playerType)
	local players = getPlayersInRoom(room.roomId)
	publishPlayerJoinRoom(room.roomId, playerId, seatId, players)
	
	return cjson.encode({roomId = room.roomId, seatId = seatId, players = players})	
    `
)
