package common

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
	"github.com/jackc/pgx"
	"github.com/nats-io/go-nats"
)

type (
	Configuration interface {
		GetDataSources() (*pgx.ConnPool, *nats.Conn)
		GetServerConfig() (string, string)
	}

	configuration struct {
		cc     *consul.Client
		prefix string
	}
)

// InitConfigService create a new config holding object
func InitConfigService(prefix string) Configuration {
	cc, e := consul.NewClient(consul.DefaultConfig())
	panicIfNil(e)
	return &configuration{cc: cc, prefix: prefix}
}

func (cfg *configuration) GetServerConfig() (string, string) {
	host, _, e := cfg.cc.KV().Get(cfg.prefix+"/server/host", nil)
	panicIfNil(e)
	port, _, e := cfg.cc.KV().Get(cfg.prefix+"/server/port", nil)
	panicIfNil(e)
	return string(host.Value), string(port.Value)
}

// GetDataSources returns the postgres and nats connections
func (cfg *configuration) GetDataSources() (*pgx.ConnPool, *nats.Conn) {
	db := cfg.connectToPostgres()
	nc := cfg.connectToNats()
	return db, nc
}

func (cfg *configuration) connectToPostgres() *pgx.ConnPool {
	connCfg := cfg.getPostgresConnectionData()
	dbPool, e := pgx.NewConnPool(pgx.ConnPoolConfig{ConnConfig: connCfg})
	panicIfNil(e)
	return dbPool
}

func (cfg *configuration) connectToNats() *nats.Conn {
	// nats://localhost:4222
	url := cfg.getNATSConnectionData()
	nc, e := nats.Connect(url)
	panicIfNil(e)
	return nc
}

func (cfg *configuration) getPostgresConnectionData() pgx.ConnConfig {
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

	cconnCfg, e := pgx.ParseDSN(dsn)
	panicIfNil(e)

	return cconnCfg
}

func (cfg *configuration) getNATSConnectionData() string {
	host, _, e := cfg.cc.KV().Get("nats/host", nil)
	panicIfNil(e)
	port, _, e := cfg.cc.KV().Get("nats/port", nil)
	panicIfNil(e)
	return fmt.Sprintf("nats://%s:%s", string(host.Value), string(port.Value))
}

func panicIfNil(e error) {
	if e != nil {
		panic(e)
	}
}
