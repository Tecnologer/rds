// Package rds provides a driver for the RDS Data API.
package rds

import (
	"database/sql"
	"database/sql/driver"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/pkg/errors"
	"github.com/tecnologer/rds/config"
)

// rdsDriver implements the driver.Driver interface.
type rdsDriver struct {
}

var _ driver.Driver = rdsDriver{}

// Open a connection. Parse the URL as a JSON config object.
func (d rdsDriver) Open(url string) (driver.Conn, error) {
	c, err := config.StringToConfig(url)
	if err != nil {
		return nil, errors.Wrap(err, "rds.driver: parse string to config")
	}
	sess, err := session.NewSession(&aws.Config{Region: aws.String(c.GetRegion())})
	if err != nil {
		return nil, errors.Wrap(err, "rds.driver.Open")
	}

	rdsAPI := rdsdataservice.New(sess)
	return &conn{
		rds:         rdsAPI,
		resourceArn: c.ResourceArn,
		secretArn:   c.SecretArn,
	}, nil
}

func init() {
	sql.Register("rds", &rdsDriver{})
}
