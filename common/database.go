package common

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var Ctx context.Context

func InitDb(loggerLevel string) {
	// Capture connection properties
	//driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	//loggerLevel := viper.GetString("logger.level")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	// create database if not exist
	temp_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port))
	if err != nil {
		panic(err)
	}
	defer temp_db.Close()
	_, err = temp_db.Exec("CREATE DATABASE IF NOT EXISTS " + database)
	if err != nil {
		panic(err)
	}
	// config
	var lvl gormlogger.LogLevel
	if loggerLevel == "Warn" {
		lvl = gormlogger.Warn
	} else {
		lvl = gormlogger.Info
	}
	newLogger := gormlogger.New(
		log.New(logger.Out, "\r\n", log.LstdFlags),
		gormlogger.Config{
			LogLevel: lvl,
		},
	)
	config := &gorm.Config{}
	config.Logger = newLogger
	// Get a database handle
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		logger.Panicln("failed to connect to database, err: " + err.Error())
	}
	// set connection pool size
	sqlDB, err := db.DB()
	if err != nil {
		logger.Panicln("failed to config db connection pool, err: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(190)
	DB = db
	logger.Info("Connected to database.")
}

func GetDB() *gorm.DB {
	return DB
}

func GetCtx() context.Context {
	return Ctx
}
