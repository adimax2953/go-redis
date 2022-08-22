package src

import (
	logtool "github.com/adimax2953/log-tool"
)

// SubscribeResult -
type SubscribeResult struct {
	Channel string
	Payload string
}

// BroadCast function - channeltype, channeltarget string, args[] string
func (s *MyScriptor) BroadCast(channeltype, channeltarget string, msg []byte) {
	strarr := []string{channeltype, channeltarget, string(msg)}
	_, err := s.Scriptor.ExecSha(BroadCastID, nil, strarr)
	if err != nil {
		logtool.LogError("BroadCast ExecSha Error", err)
	}
}

// Publish function - Channel string, data interface{} -
func (s *MyScriptor) Publish(Channel string, data interface{}) {
	err := s.Scriptor.Client.Publish(s.Scriptor.CTX, Channel, data).Err()
	if err != nil {
		logtool.LogError("Publish Error", err)
	}
}

func (s *MyScriptor) CloseSubscribe(Channel string) {

	s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel).Unsubscribe(s.Scriptor.CTX, Channel)
	s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel).Close()
}

// SubscribeString -
func (s *MyScriptor) SubscribeString(Channel string, callback func(string)) {

	go func() {
		pubsub := s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel)
		_, err := pubsub.Receive(s.Scriptor.CTX)
		if err != nil {
			logtool.LogError("SubscribeString", err.Error())
		}

		ch := pubsub.Channel()
		for message := range ch {
			payload := message.Payload
			callback(payload)
		}
	}()
}

// BroadCast - 寫入一個數字
const (
	BroadCastID       = "BroadCast"
	BroadCastTemplate = `
	local channeltype = ARGV[1]
	local channeltarget = ARGV[2]
	local payload = ARGV[3]


	redis.call("PUBLISH",channeltype,channeltarget.."~"..payload)
    `
)
