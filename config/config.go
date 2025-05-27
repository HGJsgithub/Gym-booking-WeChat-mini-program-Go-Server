package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	MySQL    `toml:"mysql"`
	Redis    `toml:"redis"`
	RabbitMQ `toml:"rabbitMQ"`
}

func (c *Config) LoadConfig() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		fmt.Println("加载配置文件失败！", err)
		panic(err)
	}
}
