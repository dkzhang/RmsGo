package login

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

func LoadLoginConfig(filepath string) (err error) {
	loginConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("LoadLoginConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(loginConfFile, &TheLoadLoginConfig)
	if err != nil {
		return fmt.Errorf("LoadLoginConfig yaml.Unmarshal error: %v", err)
	}
	return nil
}

var TheLoadLoginConfig LoginConfig

type LoginConfig struct {
	TheTokenConfig    TokenConfig    `yaml:"token"`
	ThePasswordConfig PasswordConfig `yaml:"password"`
	TheSmsConfig      SmsConfig      `yaml:"sms"`
}

type TokenConfig struct {
	Expire time.Duration `yaml:"expire"`
}

type PasswordConfig struct {
	Expire time.Duration `yaml:"expire"`
	PwType uint          `yaml:"pwType"`
	PwLen  int           `yaml:"pwLen"`
}

type SmsConfig struct {
	LockTime time.Duration `yaml:"locktime"`
}
