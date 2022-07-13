package goscriptor

import (
	"context"
	"errors"
	"sync"

	"github.com/go-redis/redis/v9"
)

// redis script defination
// the hash key definition that is used to store the script
var (
	scriptDefinition = "scriptor_v.0.0.0"
)

// Scriptor - the script manager
type Scriptor struct {
	Client                *redis.Client
	sRedisClientSyncOnce  sync.Once
	scripts               map[string]string
	redisScriptDB         int
	redisScriptDefinition string
	CTX                   context.Context
}

// New - create a new scriptor with the redis client
func New(client *redis.Client, scriptDB int, redisScriptDefinition string, scripts *map[string]string) (*Scriptor, error) {
	if client == nil {
		return nil, errors.New("'client' cannot be nil")
	}

	// new scriptor
	s := &Scriptor{}
	s.sRedisClientSyncOnce.Do(func() {
		s.Client = client
		s.scripts = make(map[string]string)
		s.redisScriptDB = scriptDB

		// if redisScriptDefinition is not empty, use the default
		if redisScriptDefinition != "" {
			s.redisScriptDefinition = redisScriptDefinition
		} else {
			s.redisScriptDefinition = scriptDefinition
		}
	})

	s.CTX = context.Background()

	// ping the redis server
	_, err := s.Client.Ping(s.CTX).Result()
	if err != nil {
		return nil, err
	}

	// load all scripts or register scripts
	scriptDescriptor, err := NewScriptDescriptor(s.CTX, s.Client, scripts, s.redisScriptDefinition, s.redisScriptDB)
	if err != nil {
		return nil, err
	}
	s.scripts = scriptDescriptor.contrainer

	return s, nil
}

// NewDB - create a new Scriptor with a new redis client
func NewDB(opt *Option, scriptDB int, redisScriptDefinition string, scripts *map[string]string) (*Scriptor, error) {
	if opt == nil {
		return nil, errors.New("'option' cannot be nil")
	}

	// new scriptor
	s := &Scriptor{}
	s.sRedisClientSyncOnce.Do(func() {
		s.Client = opt.Create()
		s.scripts = make(map[string]string)
		s.redisScriptDB = scriptDB

		// if redisScriptDefinition is not empty, use the default
		if redisScriptDefinition != "" {
			s.redisScriptDefinition = redisScriptDefinition
		} else {
			s.redisScriptDefinition = scriptDefinition
		}
	})

	s.CTX = context.Background()

	// ping the redis server
	_, err := s.Client.Ping(s.CTX).Result()
	if err != nil {
		return nil, err
	}

	// load all scripts or register scripts
	scriptDescriptor, err := NewScriptDescriptor(s.CTX, s.Client, scripts, s.redisScriptDefinition, s.redisScriptDB)
	if err != nil {
		return nil, err
	}
	s.scripts = scriptDescriptor.contrainer

	return s, nil
}

// Exec - execute the script
func (s *Scriptor) Exec(script string, keys []string, args ...interface{}) (interface{}, error) {
	if script == "" {
		return nil, errors.New("script not found")
	}

	res, err := s.Client.Eval(s.CTX, script, keys, args...).Result()

	return res, err
}

// ExecSha - execute the script
func (s *Scriptor) ExecSha(scriptname string, keys []string, args ...interface{}) (interface{}, error) {
	if s.scripts[scriptname] == "" || s.scripts == nil || len(s.scripts) == 0 {
		return nil, errors.New("script not found.")
	}
	return s.Client.EvalSha(s.CTX, s.scripts[scriptname], keys, args...).Result()
}

// stop - stop the scriptor
func (s *Scriptor) stop() {
	s.Client.Close()
}
