package config

import (
	"shop/internal/repository/mysql"

	"github.com/knadh/koanf/v2"
)

type KoanfConfig struct {
	koanf *koanf.Koanf
}

type Config struct {
	MySQL mysql.Config `koanf:"mysql"`
}



func New() KoanfConfig {
	k := koanf.New(".")
	return KoanfConfig{
		koanf: k,
	}
}

func (k KoanfConfig) GetConfig() Config {
	var cfg Config
	k.koanf.Unmarshal("", &cfg)
	return cfg
}
