package redisx

import (
	"strings"
)

type RedisxError string

func (e RedisxError) Error() string {
	return string(e)
}

const (
	InvalidCommandError        = RedisxError("redisx: invalid command")
	InvalidClusterCommandError = RedisxError("redisx: invalid cluster command")
	ClientNotFound             = RedisxError("redisx: client not found")
	PipelineIsDone             = RedisxError("redisx: pipeline is done")
	ClientAlreadyExists        = RedisxError("redisx: client already exists")
)

var clusterCmdBlackList = map[string]bool{
	"KEYS":         true,
	"MIGRATE":      true,
	"MOVE":         true,
	"OBJECT":       true,
	"RANDOMKEY":    true,
	"RENAME":       true,
	"RENAMENX":     true,
	"SCAN":         true,
	"WAIT":         true,
	"BITOP":        true,
	"MSETNX":       true,
	"BLPOP":        true,
	"BRPOP":        true,
	"BRPOPLPUSH":   true,
	"PSUBSCRIBE":   true,
	"PUBLISH":      true,
	"PUBSUB":       true,
	"PUNSUBSCRIBE": true,
	"SUBSCRIBE":    true,
	"UNSUBSCRIBE":  true,
	"EVALSHA":      true,
	"SCRIPT":       true,
	"DISCARD":      true,
	"EXEC":         true,
	"MULTI":        true,
	"UNWATCH":      true,
	"WATCH":        true,
	"CLUSTER":      true,
	"ECHO":         true,
	"QUIT":         true,
	"BGREWRITEAOF": true,
	"BGSAVE":       true,
	"CLIENT":       true,
	"COMMAND":      true,
	"CONFIG":       true,
	"DBSIZE":       true,
	"DEBUG":        true,
	"FLUSHALL":     true,
	"FLUSHDB":      true,
	"LASTSAVE":     true,
	"MONITOR":      true,
	"ROLE":         true,
	"SAVE":         true,
	"SHUTDOWN":     true,
	"SLAVEOF":      true,
	"SYNC":         true,
}

var cmdBlackList = map[string]bool{
	"KEYS":     true,
	"FLUSHALL": true,
	"FLUSHDB":  true,
	"SHUTDOWN": true,
	"CONFIG":   true,
}

func isInvalidCmd(command, mode string) error {
	command = strings.ToUpper(command)
	if _, ok := cmdBlackList[command]; ok {
		return InvalidCommandError
	}
	if mode == "cluster" {
		if _, ok := clusterCmdBlackList[command]; ok {
			return InvalidClusterCommandError
		}
	}
	return nil
}
