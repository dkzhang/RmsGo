package logMap

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// TODO: cancel log variable as global variable.
var theLogFileConfig LogFileConfig

type LogFileConfig struct {
	LogFile map[string]string `yaml:"logfile,omitempty"`
}

func LoadLogConfig(filepath string) (err error) {

	logConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("LoadLogConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(logConfFile, &theLogFileConfig)
	if err != nil {
		return fmt.Errorf("LoadLogConfig yaml.Unmarshal error: %v", err)
	}
	return nil
}
