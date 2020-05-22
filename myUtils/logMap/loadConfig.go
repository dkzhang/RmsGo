package logMap

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func init() {

}

var theLogFileConfig LogFileConfig

type LogFileConfig struct {
	LogFile map[string]string `yaml:"logfile,omitempty"`
}

func LoadLogConfig(filepath string) (theLogConfig *LogFileConfig, err error) {
	theLogConfig = &LogFileConfig{}
	logConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("LoadLogConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(logConfFile, theLogConfig)
	if err != nil {
		return nil, fmt.Errorf("LoadLogConfig yaml.Unmarshal error: %v", err)
	}
	return theLogConfig, nil
}
