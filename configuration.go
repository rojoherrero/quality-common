package common

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

const (
	emptyString = ""
)

type (
	Configuration interface {
		GetNatsDSN() (string, error)
		GetPostgresDSN() (string, error)
		GetServerConfig() (string, string, error)
	}

	configuration struct {
		cc     *consul.Client
		prefix string
	}
)

// InitConfigService create a new config holding object
func InitConfigService(prefix string) (Configuration, error) {
	cc, e := consul.NewClient(consul.DefaultConfig())
	if e != nil {
		return nil, e
	}
	return &configuration{cc: cc, prefix: prefix}, nil
}

func (cfg *configuration) GetServerConfig() (string, string, error) {
	host, _, e := cfg.cc.KV().Get(cfg.prefix+"/server/host", nil)
	if e != nil {
		return emptyString, emptyString, e
	}
	port, _, e := cfg.cc.KV().Get(cfg.prefix+"/server/port", nil)
	if e != nil {
		return emptyString, emptyString, e
	}
	return string(host.Value), string(port.Value), nil
}

func (cfg *configuration) GetPostgresDSN() (string, error) {
	host, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/host", nil)
	if e != nil {
		return emptyString, e
	}
	port, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/port", nil)
	if e != nil {
		return emptyString, e
	}
	user, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/user", nil)
	if e != nil {
		return emptyString, e
	}
	pwd, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/pwd", nil)
	if e != nil {
		return emptyString, e
	}
	dBName, _, e := cfg.cc.KV().Get(cfg.prefix+"/pg/dbName", nil)
	if e != nil {
		return emptyString, e
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		string(host.Value), string(port.Value), string(user.Value), string(pwd.Value), string(dBName.Value))

	return dsn, nil
}

func (cfg *configuration) GetNatsDSN() (string, error) {
	host, _, e := cfg.cc.KV().Get("nats/host", nil)
	if e != nil {
		return emptyString, e
	}
	port, _, e := cfg.cc.KV().Get("nats/port", nil)
	if e != nil {
		return emptyString, e
	}
	return fmt.Sprintf("nats://%s:%s", string(host.Value), string(port.Value)), nil
}

