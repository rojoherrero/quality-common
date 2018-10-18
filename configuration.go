package common

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Configuration struct {
	Messaging *Messaging
	Database  *Database
	Server    *Server
	Logging   *Logging
}

type Server struct {
	Port string
}

type Messaging struct {
	Host string
	Port string
}

type Database struct {
	Username string
	Password string
	Host     string
	Port     uint16
	DbName   string
}

type Logging struct {
	Level string
}

func GetConfiguration(name, path string) *Configuration {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(name)
	v.AddConfigPath(path)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	config := new(Configuration)
	if err := v.Unmarshal(config); err != nil {
		log.Fatal(err)
	}
	return config
}

func (cfg *Configuration) ConnectToPostgres() *sqlx.DB {
	// username:password@protocol(address)/dbname?param=value
	pqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DbName)

	var (
		db *sqlx.DB
		e error
	)

	if db, e = sqlx.Connect("postgres", pqInfo); e != nil {
		panic(e)
	}
	return db
}

func (cfg *Configuration) ConnectToNats() *nats.Conn {
	// nats://localhost:4222
	url := fmt.Sprintf("nats://%s:%s", cfg.Messaging.Host, cfg.Messaging.Port)

	var(
		nc *nats.Conn
		e error
	)

	if nc, e = nats.Connect(url); e != nil {
		panic(e)
	}
	return nc
}
