package config

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	configFileEnvVar = "HOG_CONFIG_PATH"
)

func LoadConfig() (*HOGConfig, error) {
	fileName := os.Getenv(configFileEnvVar)

	logrus.Infof("opening config file: '%s'", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer file.Close()

	var result HOGConfig

	err = json.NewDecoder(file).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("parsing JSON config file: %w", err)
	}

	return &result, nil
}
