package security

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/genPasswd"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type LoginSecurity struct {
	TheTokenSecurity TokenSecurity `yaml:"TokenSecurity"`
}

type TokenSecurity struct {
	Key string `yaml:"Key"`
}

func LoadLoginSecurity(filepath string) (sec LoginSecurity, err error) {
	loginSecFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return LoginSecurity{}, fmt.Errorf("LoadLoginSecurity ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(loginSecFile, &sec)
	if err != nil {
		return LoginSecurity{}, fmt.Errorf("LoadLoginSecurity yaml.Unmarshal error: %v", err)
	}
	return sec, nil
}

var seedOnce sync.Once

func GenLoginSecurity() (sec LoginSecurity, err error) {
	pwLen := 12
	pwType := genPasswd.FlagLowerChar
	seedOnce.Do(func() {
		genPasswd.RandomSeed()
	})
	return LoginSecurity{
		TheTokenSecurity: TokenSecurity{
			Key: genPasswd.GeneratePasswd(pwLen, pwType)}}, nil
}
