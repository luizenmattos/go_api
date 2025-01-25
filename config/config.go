package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
	env    string
)

func Init() error {
	//setting eviorment
	env = "dev"

	//Initializar SQLite
	var err error
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)
	return logger
}

func GetEnv() string {
	return env
}
