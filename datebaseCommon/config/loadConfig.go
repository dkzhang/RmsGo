package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//load database configuration from file. Including PostgreSQL and Redis.
type DbConfig struct {
	ThePgConfig    PgConfig    `yaml:"PostgreSQL"`
	TheRedisConfig RedisConfig `yaml:"Redis"`
}

type PgConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`

	//"host=ras-pg user=postgres password=%s dbname=ras sslmode=disable"
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
}

func LoadDbConfig(filepath string) (theDbConfig *DbConfig, err error) {
	theDbConfig = &DbConfig{}
	dbConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("LoadDbConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(dbConfFile, theDbConfig)
	if err != nil {
		return nil, fmt.Errorf("LoadDbConfig yaml.Unmarshal error: %v", err)
	}
	return theDbConfig, nil
}
