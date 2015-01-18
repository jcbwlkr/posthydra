package posthydra

import "code.google.com/p/gcfg"

type Config struct {
	WildApricot struct {
		Key       string
		AccountId string
	}
}

func NewConfig(path string) *Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, path)
	if err != nil {
		panic(err)
	}

	return &cfg
}
