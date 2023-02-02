package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewSQLDb(config SQLConfig) {

	fmt.Printf(
		"user=%s"+
			"password=%s"+
			"dbname=%s",
		config.User,
		config.Password,
		config.DBName,
	)

	sqlx.Connect(config.Driver, config.DBName)
}
