package config_reader

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v2"
)

type YAMLConfig struct {
}

func (yc *YAMLConfig) GetGroupConfig(idx int) (*GroupConfig, error) {
	filename := filepath.Join("configs", strconv.Itoa(idx)+".yaml")

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
		return nil, err
	}

	cfg := GroupConfig{}
	err = yaml.Unmarshal(content, &cfg)

	if err != nil {
		log.Printf("ERROR: %s\n", err.Error())
	}

	return &cfg, nil
}

func NewYAMLConfig() (Config, error) {
	return &YAMLConfig{}, nil
}
