package posthydra

import "code.google.com/p/gcfg"

type Config struct {
	WildApricot WildApricotConfig
}

func NewConfig(path string) (*Config, error) {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, path)

	return &cfg, err
}
