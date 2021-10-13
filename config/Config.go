package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	//More info here: https://docs.aws.amazon.com/secretsmanager/latest/userguide/tutorials_basic.html
	ResourceArn string
	//More info here: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Getting
	SecretArn string
	region    string `json:"-"`
}

//String returns Config as JSON string
func (c *Config) String() string {
	str, e := json.Marshal(c)
	if e != nil {
		return fmt.Sprintf("error: %v", e)
	}

	return string(str)
}

//GetRegion extracts the region from the resouce ARN
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

//GetDefaultConfig a *Config with the data from Env Vars (RDS_SECRET_ARN & RDS_RESOURCE_ARN)
func GetDefaultConfig() *Config {
	return &Config{
		SecretArn:   os.Getenv("RDS_SECRET_ARN"),
		ResourceArn: os.Getenv("RDS_RESOURCE_ARN"),
	}
}

func StringToConfig(cfg string) (c *Config, err error) {
	if err = json.Unmarshal([]byte(cfg), &c); err != nil {
		return nil, errors.Wrap(err, "rds.config: parse string to config")
	}

	return c, nil
}
