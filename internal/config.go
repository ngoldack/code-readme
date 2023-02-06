package internal

import "github.com/spf13/viper"

type Config struct {
	InputFile  string `yaml:"inputFile"`
	OutputFile string `yaml:"outputFile"`

	Born   string       `yaml:"born"`
	Github GithubConfig `yaml:"github"`
}

type GithubConfig struct {
	Token    string `yaml:"token"`
	Username string `yaml:"username"`

	Repositories []string `yaml:"repositories"`
}

func ParseConfig() (*Config, error) {
	cfg := &Config{}
	err := viper.Unmarshal(cfg)
	return cfg, err
}
