package Config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var config ConfigStruct

func GetConfig() ConfigStruct {
	return config
}

func LoadConfig(configFile string) error {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}
