package cmd

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Aliyun *AliConfig `yaml:"aliyun"`
}

type AliConfig struct {
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
}

func GetConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}