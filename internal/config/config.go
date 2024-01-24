package config

import (
	"github.com/iot-synergy/synergy-common/config"
	"github.com/iot-synergy/synergy-common/plugins/mq/asynq"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	RedisConf    config.RedisConf
	AsynqConf    asynq.AsynqConf
	TaskConf     TaskConf
}

type TaskConf struct {
	EnableScheduledTask bool `json:",default=true"`
	EnableDPTask        bool `json:",default=true"`
}
