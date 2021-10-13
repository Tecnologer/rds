package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	ResourceArn string
	SecretArn   string
	region      string `json:"-"`
}

func (c *Config) String() string {
	str, e := json.Marshal(c)
	if e != nil {
		return fmt.Sprintf("error: %v", e)
	}

	return string(str)
}

func (c *Config) GetRegion() string {
	if c.region == "" {
		tokens := strings.Split(c.ResourceArn, ":")
		if len(tokens) < 4 {
			return ""
		}
		c.region = tokens[3]
	}
	return c.region
}
