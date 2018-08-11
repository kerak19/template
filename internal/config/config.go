package config

import (
	"time"

	"github.com/BurntSushi/toml"
)

// Database contains database informations
type Database struct {
	Migrations string `toml:"migrations"`
	URL        string `toml:"url"`
}

// Config contains all application configurable data
type Config struct {
	Addr                  string        `toml:"addr"`
	DebugMode             bool          `toml:"debugMode"`
	JWTSecret             string        `toml:"jwt_secret"`
	ServerShutdownTimeout time.Duration `toml:"server_shutdown_timeout"`

	Database Database `toml:"database"`
}

// ReadFromFiles reads TOML configuration from multiple file sources. Each next
// file can overwrite previous ones if fields are repeated
func ReadFromFiles(paths ...string) (c Config, err error) {
	for _, v := range paths {
		_, err = toml.DecodeFile(v, &c)
		if err != nil {
			return
		}
	}
	return
}
