package cache_redis_config

import (
	"fmt"
	"strings"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewConfig(host string, port int, password string, db int) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		Password: password,
		DB:       db,
	}
}

func (c *Config) Validate() error {
	if c.Host == "" {
		return fmt.Errorf("host is required")
	}
	if c.Port == 0 {
		return fmt.Errorf("port is required")
	}
	if c.Password == "" {
		return fmt.Errorf("password is required")
	}
	if c.DB == 0 {
		return fmt.Errorf("database number is required")
	}
	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("redis://%s:%d@%s/%d", c.Password, c.Port, c.Host, c.DB)
}

func ParseConfig(str string) (*Config, error) {
	parts := strings.SplitN(str, "://", 2)
	if len(parts)!= 2 {
		return nil, fmt.Errorf("invalid config format")
	}
	var password, host, port, db string
	if strings.HasPrefix(parts[1], "@") {
		password = parts[1][1:]
		host = strings.SplitN(password, ":", 2)[0]
		port = strings.SplitN(password, ":", 1)[1]
	} else {
		host = parts[1]
	}
	db = strings.SplitN(host, "/", 2)[1]
	return NewConfig(host, 6379, password, 0), nil
}