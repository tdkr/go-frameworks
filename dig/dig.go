package main

import (
	"encoding/json"
	"go.uber.org/dig"
	"log"
	"os"
)

type Config struct {
	Prefix string
}

func main() {
	c := dig.New()

	err := c.Provide(func() (*Config, error) {
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		panic(err)
	}

	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		panic(err)
	}

	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		panic(err)
	}
}
