package common

import (
	"fmt"
	"log"

	consul "github.com/hashicorp/consul/api"
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
)

type Configuration struct {
	cc     *consul.Client
	prefix string
}

type Server struct {
	Host string
	Port string
}

type Logging struct {
	Level string
}

func InitConfigService(prefix string) *Configuration {
	cc, e := consul.NewClient(consul.DefaultConfig())
	panicIfNil(e)
	return &Configuration{cc: cc, prefix: prefix}
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
func (cfg *Configuration) GetDataSources() (*pgx.ConnPool, *nats.Conn) {
	db := cfg.connectToPostgres()
	nc := cfg.connectToNats()
	return db, nc
}

func (cfg *Configuration) connectToPostgres() *pgx.ConnPool {
	// username:password@protocol(address)/dbname?param=value
	connCfg := cfg.getPostgresConnectionData()
	db, e := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connCfg})
	panicIfNil(e)
	return db
}

func (cfg *Configuration) connectToNats() *nats.Conn {
	// nats://localhost:4222
	url := cfg.getNATSConnectionData()
	nc, e := nats.Connect(url)
	panicIfNil(e)
	return nc
}

func (cfg *Configuration) getPostgresConnectionData() pgx.ConnConfig {
	host, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/host", nil)
	panicIfNil(e)
	port, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/port", nil)
	panicIfNil(e)
	user, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/user", nil)
	panicIfNil(e)
	pwd, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/pwd", nil)
	panicIfNil(e)
	dBName, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/dbName", nil)
	panicIfNil(e)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		string(host.Value), string(port.Value), string(user.Value), string(pwd.Value), string(dBName.Value))

	config, e := pgx.ParseDSN(dsn)
	panicIfNil(e)

	return config

}

func (cfg *Configuration) getNATSConnectionData() string {
	host, _, _ := cfg.cc.KV().Get("nats/host", nil)
	port, _, _ := cfg.cc.KV().Get("nats/port", nil)
	return fmt.Sprintf("nats://%s:%s", string(host.Value), string(port.Value))
}

func panicIfNil(e error) {
	if e != nil {
		panic(e)
	}
}
