package shortMessageService

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SMSConfig struct {
	ID         string `yaml:"id"`
	Key        string `yaml:"key"`
	AppID      string `yaml:"appid"`
	Sign       string `yaml:"sign"`
	TemplateID string `yaml:"templateid"`
}

var TheSMSConfig SMSConfig

func LoadSmsConfig(filepath string) (err error) {
	smsConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("LoadSMSConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(smsConfFile, &TheSMSConfig)
	if err != nil {
		return fmt.Errorf("LoadSMSConfig yaml.Unmarshal error: %v", err)
	}
	return nil
}
