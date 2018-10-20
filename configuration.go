package common

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"

	consul "github.com/hashicorp/consul/api"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server  *Server
	Logging *Logging
}

type Server struct {
	Port string
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

// GetDataSources returns the postgres and nats connections
func GetDataSources(prefix string) (*pgx.ConnPool, *nats.Conn) {
	kv := getKVClient()
	db := connectToPostgres(prefix, kv)
	nc := connectToNats(kv)
	return db, nc
}

func getKVClient() *consul.KV {
	client, e := consul.NewClient(consul.DefaultConfig())
	if e != nil {
		panic(e)
	}
	kv := client.KV()
	return kv
}

func connectToPostgres(prefix string, kv *consul.KV) *pgx.ConnPool {
	// username:password@protocol(address)/dbname?param=value
	connCfg := getPostgresConnectionData(prefix, kv)

	var db *pgx.ConnPool
	var e error

	if db, e = pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connCfg}); e != nil {
		panic(e)
	}
	return db
}

func connectToNats(kv *consul.KV) *nats.Conn {
	// nats://localhost:4222
	url := getNATSConnectionData(kv)
	var nc *nats.Conn
	var e error

	if nc, e = nats.Connect(url); e != nil {
		panic(e)
	}
	return nc
}

func getPostgresConnectionData(prefix string, kv *consul.KV) pgx.ConnConfig {
	host, _, _ := kv.Get(prefix+"/pg/host", nil)
	port, _, _ := kv.Get(prefix+"/pg/portT", nil)
	user, _, _ := kv.Get(prefix+"/pg/user", nil)
	pwd, _, _ := kv.Get(prefix+"/pg/pwd", nil)
	dBName, _, _ := kv.Get(prefix+"/pg/dbName", nil)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		string(host.Value), string(port.Value), string(user.Value), string(pwd.Value), string(dBName.Value))

	var config pgx.ConnConfig
	var e error

	if config, e = pgx.ParseDSN(dsn); e != nil {
		panic(e)
	}

	return config

}

func getNATSConnectionData(kv *consul.KV) string {
	host, _, _ := kv.Get("nats/host", nil)
	port, _, _ := kv.Get("nats/port", nil)

	return fmt.Sprintf("nats://%s:%s", string(host.Value), string(port.Value))
}
