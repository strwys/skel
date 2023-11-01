package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func (cfg Config) NewDatabase() (*sql.DB, error) {

	dbConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.User,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Name,
	)

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	maxOpenConns := viper.GetInt("DBMAXOPENCONNS")
	if maxOpenConns == 0 {
		maxOpenConns = 25
	}
	db.SetMaxOpenConns(maxOpenConns)

	maxIdleConns := viper.GetInt("DBMAXIDLECONNS")
	if maxIdleConns == 0 {
		maxIdleConns = 10
	}
	db.SetMaxIdleConns(maxIdleConns)

	maxLifetime := viper.GetInt("DBMAXLIFETIME")
	if maxLifetime == 0 {
		maxLifetime = 300
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	return db, nil
}
