package config

import (
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
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`

	//"host=ras-pg user=postgres password=%s dbname=ras sslmode=disable"
}

type RedisConfig struct {
}

var TheDbConfig DbConfig

func LoadDbConfig(filepath string) (err error) {
	dbConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dbConfFile, &TheDbConfig)
	if err != nil {
		return err
	}
	return nil
}
