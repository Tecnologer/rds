package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tecnologer/rds"
	"github.com/tecnologer/rds/config"
)

func main() {
	config := &config.Config{
		SecretArn:   os.Getenv("RDS_SECRET_ARN"),
		ResourceArn: os.Getenv("RDS_RESOURCE_ARN"),
	}

	db, err := sql.Open("rds", config.String())
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("connected")

	defer db.Close()

}
