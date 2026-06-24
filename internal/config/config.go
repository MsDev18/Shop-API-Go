package config

import (
	"fmt"

	"github.com/knadh/koanf/v2"
)

type KoanfConfig struct {
	koanf *koanf.Koanf
}

type Config struct {

}

func New() KoanfConfig {
	k := koanf.New(".")
	return KoanfConfig{
		koanf: k,
	}
}

func (k KoanfConfig) GetConfig () Config {
	var cfg Config
	k.koanf.Unmarshal("" , &cfg)
	fmt.Println(cfg)
	return cfg
}

