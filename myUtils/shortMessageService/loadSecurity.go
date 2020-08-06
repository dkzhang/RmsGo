package shortMessageService

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SmsSecurity struct {
	ID               string `yaml:"id"`
	Key              string `yaml:"key"`
	AppID            string `yaml:"appid"`
	Sign             string `yaml:"sign"`
	PwdTemplateID    string `yaml:"pwd_templateid"`
	NotifyTemplateID string `yaml:"notify_templateid"`
}

func LoadSmsSecurity(filepath string) (theSmsSecurity SmsSecurity, err error) {
	smsConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return SmsSecurity{}, fmt.Errorf("LoadSMSConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(smsConfFile, &theSmsSecurity)
	if err != nil {
		return SmsSecurity{}, fmt.Errorf("LoadSMSConfig yaml.Unmarshal error: %v", err)
	}
	return theSmsSecurity, nil
}
