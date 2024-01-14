package config

import (
	"time"
)

type Config struct {
	Timeout       time.Duration
	UseAllNumbers bool
	UseBrackets   bool
	Debug         bool
}

var Default = Config{
	Timeout:       28 * time.Second,
	UseAllNumbers: true,
	Debug:         false,
}
