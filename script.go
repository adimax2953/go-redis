package go-redis

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v9"
)

// Load the lua script sha
// Store the SHA1 and script name in redis
var (
	loadLuaScriptTemplate = `
		redis.pcall('SELECT', ARGV[1])
		return redis.call('HGETALL', KEYS[1])
	`

	getLuaScriptTemplate = `
		redis.pcall('SELECT', ARGV[1])
		return redis.call('HGET', KEYS[1], ARGV[2])
	`

	setLuaScriptTemplate = `
		redis.pcall('SELECT', ARGV[1])
		return redis.call('HSET', KEYS[1], ARGV[2], ARGV[3])
	`

	existsLuaScriptTemplate = `
		redis.pcall('SELECT', ARGV[1])
		return redis.call('EXISTS', KEYS[1])
	`

	hexistsLuaScriptTemplate = `
		redis.pcall('SELECT', ARGV[1])
		return redis.call('HEXISTS', KEYS[1], ARGV[2])
	`
)

// Script is a script descriptor
type ScriptDescriptor struct {
	contrainer map[string]string
	Scripts    map[string]string
}

// NewScriptDescriptor creates a new script descriptor
func NewScriptDescriptor(ctx context.Context, client *redis.Client, scripts *map[string]string, redisScriptDefinition string, db int) (*ScriptDescriptor, error) {
	if client == nil {
		return nil, errors.New("'client' is invalid")
	}

	// Create a new script descriptor
	scriptDescriptor := &ScriptDescriptor{}

	if scripts == nil || len(*scripts) == 0 {
		// Lodad the lua script sha1
		err := scriptDescriptor.LoadScripts(ctx, client, redisScriptDefinition, db)
		if err != nil {
			return nil, err
		}
		return scriptDescriptor, nil
	}

	// Load the script
	// Load a script into the scripts cache, without executing it.
	err := scriptDescriptor.Register(ctx, client, scripts, redisScriptDefinition, db)
	if err != nil {
		return nil, err
	}

	return scriptDescriptor, nil
}

// Registers a script
func (scriptDescriptor *ScriptDescriptor) Register(ctx context.Context, client *redis.Client, scripts *map[string]string, redisScriptDefinition string, db int) error {
	scriptDescriptor.contrainer = make(map[string]string)

	// Registers a script
	for name, body := range *scripts {
		var err error
		var sha1 string

		// if the script key & mkey is exists
		// get the script sha1 and ScriptExists the sha1
		// countinue the script
		sha1, err = availableLuaScript(ctx, client, redisScriptDefinition, db, name)
		if err == nil {
			scriptDescriptor.contrainer[name] = sha1
			continue
		}

		sha1, err = client.ScriptLoad(ctx, body).Result()

		if err != nil {
			return err
		}

		err = setLuaScript(ctx, client, redisScriptDefinition, name, sha1, db)
		if err != nil {
			return err
		}

		scriptDescriptor.contrainer[name] = sha1
	}

	return nil
}

// LoadScripts loads the scripts
func (scriptDescriptor *ScriptDescriptor) LoadScripts(ctx context.Context, client *redis.Client, redisScriptDefinition string, db int) error {
	if client == nil {
		return errors.New("'client' can not be nil.")
	}

	// // check if the script key exists
	// err := keyExistsLuaScript(ctx, client, redisScriptDefinition, db)
	// if err != nil {
	// 	return err
	// }

	// Load the script
	res, err := loadLuaScript(ctx, client, redisScriptDefinition, db)
	if err != nil {
		return err
	}

	if v, ok := res.([]interface{}); ok {
		count := len(v)
		if count == 0 {
			return nil
		}

		// Parse the script name and sha1
		// checking the existence of the scripts in the script cache.
		scriptDescriptor.contrainer = make(map[string]string)
		for i := 0; i < count; i = i + 2 {
			key, value := v[i], v[i+1]

			err := scriptExists(ctx, client, value.(string))
			if err != nil {
				return err
			}
			scriptDescriptor.contrainer[key.(string)] = value.(string)
		}
	}
	return nil
}

// keyExistsLuaScript - check if the script key exists
func keyExistsLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, db int) error {
	// check if the script key exists
	exists, err := client.Eval(ctx, existsLuaScriptTemplate, []string{redisScriptDefinition}, db).Result()
	if err != nil {
		return err
	}
	if exists.(int64) == 0 {
		return errors.New("Script key does not exist.")
	}

	return nil
}

// mkeyExistsLuaScript - check if the script member key exists
func mkeyExistsLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, mkey string, db int) error {
	// check if the script member key exists
	exists, err := client.Eval(ctx, hexistsLuaScriptTemplate, []string{redisScriptDefinition, mkey}, db).Result()
	if err != nil {
		return err
	}
	if exists.(int64) == 0 {
		return errors.New("Script key does not exist.")
	}

	return nil
}

// getLuaScript - get the lua script sha1
func getLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, name string, db int) (string, error) {

	exists, err := client.Eval(ctx, getLuaScriptTemplate, []string{redisScriptDefinition}, db, name).Result()
	if err != nil {
		return "", err
	}
	if exists.(string) == "" {
		return "", errors.New("Script is not already")
	}

	return exists.(string), nil
}

// setLuaScript - set the lua script sha1
func setLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, name string, sha1 string, db int) error {
	_, err := client.Eval(ctx, setLuaScriptTemplate, []string{redisScriptDefinition}, db, name, sha1).Result()
	if err != nil {
		return err
	}

	return nil
}

// loadLuaScript - Load the script
func loadLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, db int) (interface{}, error) {
	res, err := client.Eval(ctx, loadLuaScriptTemplate, []string{redisScriptDefinition}, db).Result()
	return res, err
}

// scriptExists - checking the existence of the scripts in the script cache.
func scriptExists(ctx context.Context, client *redis.Client, sha1 string) error {
	exists, err := client.ScriptExists(ctx, sha1).Result()
	if err != nil {
		return err
	}
	if !exists[0] {
		return errors.New("script is not exists. please reload your script")
	}
	return nil
}

// availableLuaScript - check if the script is available
func availableLuaScript(ctx context.Context, client *redis.Client, redisScriptDefinition string, db int, name string) (string, error) {
	var err error
	var sha1 string
	// check if the script key & mkey is exists
	// get the script sha1
	err = keyExistsLuaScript(ctx, client, redisScriptDefinition, db)
	if err != nil {
		return "", err
	}
	err = mkeyExistsLuaScript(ctx, client, redisScriptDefinition, name, db)
	if err != nil {
		return "", err
	}
	sha1, err = getLuaScript(ctx, client, redisScriptDefinition, name, db)
	if err != nil {
		return "", err
	}

	err = scriptExists(ctx, client, sha1)
	if err != nil {
		return "", err
	}

	return sha1, nil
}
