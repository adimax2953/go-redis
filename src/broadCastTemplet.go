package src

import (
	"sync"

	logtool "github.com/adimax2953/log-tool"
)

// SubscribeResult -
type SubscribeResult struct {
	Channel string
	Payload string
}

// Publish function - Channel string, data interface{} -
func (s *MyScriptor) Publish(Channel string, data interface{}) {
	s.Scriptor.Client.Publish(s.Scriptor.CTX, Channel, data)
}

func (s *MyScriptor) CloseSubscribe(Channel string) {

	s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel).Unsubscribe(s.Scriptor.CTX, Channel)
	s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel).Close()
}

var Pubsub sync.Map

// SubscribeString -
func (s *MyScriptor) SubscribeString(Channel string, callback func(string)) {

	go func() {
		pubsub := s.Scriptor.Client.Subscribe(s.Scriptor.CTX, Channel)
		_, err := pubsub.Receive(s.Scriptor.CTX)
		if err != nil {
			logtool.LogFatal("SubscribeString", err.Error())
		}

		Pubsub.Store(Channel, &pubsub)

		ch := pubsub.Channel()

		for message := range ch {
			payload := message.Payload
			callback(payload)
		}
	}()
}

// Publish - 寫入一個數字
const (
	PublishID       = "Publish"
	PublishTemplate = `
	local channeltype = ARGV[1]
	local channeltarget = ARGV[2]
	local payload = ARGV[3]


	redis.call("PUBLISH",channeltype,channeltarget.."~"..payload)
    `
)
