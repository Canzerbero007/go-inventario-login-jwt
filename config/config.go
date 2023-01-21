package config

// import lib go
import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Config struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*Config, error) {
	var c Config

	err := yaml.Unmarshal(configFile, &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
