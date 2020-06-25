package common

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Environment string

const (
	Production  Environment = "prd"
	Staging     Environment = "stg"
	Development Environment = "dev"
)

func (e *Environment) Decode(value string) error {
	switch Environment(value) {
	case Production:
		*e = Production
	case Staging:
		*e = Staging
	case Development:
		*e = Development
	default:
		return fmt.Errorf("invalid Environment")
	}
	return nil
}

type Config struct {
	App App
	GAE GAE
	K   K
}

func (c *Config) IsGAE() bool {
	return c.GAE.Service != ""
}

func (c *Config) IsRun() bool {
	return c.K.Service != ""
}

type App struct {
	Name string      `required:"true"`
	Env  Environment `required:"true"`
	Hoge string
}

type GAE struct {
	Application  string
	DeploymentID string `split_words:"true"`
	Env          string
	Instance     string
	MemoryMB     int `split_words:"true"`
	Runtime      string
	Service      string
	Version      string
}

type K struct {
	Service       string
	Revision      string
	Configuration string
}

func LoadConfig() (c *Config, err error) {
	var app App
	err = envconfig.Process("app", &app)
	if err != nil {
		return
	}

	var gae GAE
	err = envconfig.Process("gae", &gae)
	if err != nil {
		return
	}

	var k K
	err = envconfig.Process("k", &k)
	if err != nil {
		return
	}

	return &Config{
		App: app,
		GAE: gae,
		K:   k,
	}, nil
}
