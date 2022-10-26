package config_reader

import "fmt"

func NewConfig(s string) (Config, error) {
	switch s {
	case "static":
		return NewStaticConfig()
	case "yaml":
		return NewYAMLConfig()
	}

	return nil, fmt.Errorf("unknown ConfigType: %s", s)
}
