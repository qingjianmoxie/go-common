package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadYaml(filename string, config interface{}) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, config)
	return err
}