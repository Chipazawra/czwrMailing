package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/Chipazawra/czwrmailing/internal/services/auth"
	"gopkg.in/yaml.v2"
)

// Config struct for webapp config
type Config struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`
		// Port is the local machine TCP Port to bind the HTTP Server to
		Port string `yaml:"port"`
		// true - enable logging, false - disble logging
		Log bool `yaml:"log"`
		// true - logging to file, false  - to stdOut
		LogToFile bool `yaml:"logtofile"`
		// if LogToFile = true, path to log file
		LogPath string `yaml:"logpath"`
	}
	auth.AuthConf
}

// NewConfig returns a new decoded Config struct
func newConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func parseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func Load() (*Config, error) {

	cfgPath, err := parseFlags()
	if err != nil {
		return nil, err
	}
	cfg, err := newConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
