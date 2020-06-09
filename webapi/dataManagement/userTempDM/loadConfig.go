package userTempDM

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

//load database configuration from file. Including PostgreSQL and Redis.
type LoginConfig struct {
	TheTokenConfig    TokenConfig    `yaml:"token"`
	ThePasswordConfig PasswordConfig `yaml:"password"`
	TheSmsConfig      SmsConfig      `yaml:"sms"`
}

type TokenConfig struct {
	ExpireStr string        `yaml:"expire"`
	Expire    time.Duration `yaml:"-"`
}

type PasswordConfig struct {
	ExpireStr string        `yaml:"expire"`
	Expire    time.Duration `yaml:"-"`
	PwType    int           `yaml:"pwType"`
	PwLen     int           `yaml:"pwLen"`
}

type SmsConfig struct {
	LockTimeStr string        `yaml:"locktime"`
	LockTime    time.Duration `yaml:"-"`
}

func LoadLoginConfig(filepath string) (cfg LoginConfig, err error) {
	dbConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return LoginConfig{}, fmt.Errorf("LoadLoginConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(dbConfFile, &cfg)
	if err != nil {
		return LoginConfig{}, fmt.Errorf("LoadLoginConfig yaml.Unmarshal error: %v", err)
	}

	cfg.TheTokenConfig.Expire, err = time.ParseDuration(cfg.TheTokenConfig.ExpireStr)
	if err != nil {
		return LoginConfig{},
			fmt.Errorf("time.ParseDuration(cfg.TheTokenConfig.ExpireStr) error: %v", err)
	}
	cfg.ThePasswordConfig.Expire, err = time.ParseDuration(cfg.ThePasswordConfig.ExpireStr)
	if err != nil {
		return LoginConfig{},
			fmt.Errorf("time.ParseDuration(cfg.ThePasswordConfig.ExpireStr) error: %v", err)
	}
	cfg.TheSmsConfig.LockTime, err = time.ParseDuration(cfg.TheSmsConfig.LockTimeStr)
	if err != nil {
		return LoginConfig{},
			fmt.Errorf("time.ParseDuration(cfg.TheSmsConfig.LockTimeStr) error: %v", err)
	}

	return cfg, nil
}
