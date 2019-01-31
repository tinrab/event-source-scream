package config

import (
	"github.com/tinrab/kit/cfg"
)

type Config struct {
	Database DatabaseConfig `cfg:"database"`
	Bus      BusConfig      `cfg:"bus"`
	API      APIConfig      `cfg:"api"`
}

type DatabaseConfig struct {
	Host     string `cfg:"host"`
	Port     uint16 `cfg:"port"`
	Name     string `cfg:"name"`
	User     string `cfg:"user"`
	Password string `cfg:"password"`
}

type BusConfig struct {
	Host string `cfg:"host"`
	Port uint16 `cfg:"port"`
}

type APIConfig struct {
	Port uint16 `cfg:"port"`
}

func Load(filename string) (*Config, error) {
	data := cfg.New()

	if err := data.LoadFile(filename); err != nil {
		return nil, err
	}

	c := &Config{}
	if err := data.Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}
