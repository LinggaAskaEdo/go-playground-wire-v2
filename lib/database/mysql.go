package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func InitMySQL(ctx context.Context) (*sql.DB, error) {
	config := ctx.Value("mysql").(common.MySQLConfiguration)

	databaseURI := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
		config.MySQLName,
	)

	db, err := sql.Open("mysql", databaseURI)
	if err != nil {
		return db, err
	}

	db.SetMaxIdleConns(config.MySQLMaxIdleConns)
	db.SetMaxOpenConns(config.MySQLMaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(config.MySQLConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(config.MySQLConnMaxIdleTime) * time.Minute)

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
