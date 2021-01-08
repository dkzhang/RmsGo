package security

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//load database configuration from file. Including PostgreSQL and Redis.
type DbSecurity struct {
	ThePgSecurity    PgSecurity    `yaml:"PostgreSQL"`
	TheRedisSecurity RedisSecurity `yaml:"Redis"`
}

type PgSecurity struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Sslmode  string `yaml:"sslmode"`

	//"host=ras-pg user=postgres password=%s dbname=ras sslmode=disable"
}

type RedisSecurity struct {
	Host string `yaml:"host"`
	//Password string `yaml:"password"`
}

func LoadDbSecurity(filepath string) (dbSec DbSecurity, err error) {
	dbSecFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return DbSecurity{}, fmt.Errorf("LoadDbSecurity ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(dbSecFile, &dbSec)
	if err != nil {
		return DbSecurity{}, fmt.Errorf("LoadDbSecurity yaml.Unmarshal error: %v", err)
	}
	return dbSec, nil
}
