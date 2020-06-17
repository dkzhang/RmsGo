package logMap

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadLogConfig(filepath string) (theLogMap LogMap, err error) {
	logConfFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return LogMap{}, fmt.Errorf("LoadLogConfig ioutil.ReadFile error: %v", err)
	}
	err = yaml.Unmarshal(logConfFile, &theLogMap)
	if err != nil {
		return LogMap{}, fmt.Errorf("LoadLogConfig yaml.Unmarshal error: %v", err)
	}
	return theLogMap, nil
}
