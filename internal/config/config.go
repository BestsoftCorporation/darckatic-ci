package config

import (
	"gopkg.in/yaml.v3"
	"log"
)

// Environment represents the deployment environment.
type Environment struct {
	Name          string `yaml:"name"`
	ServerAddress string `yaml:"server_address"`
}

// GitCredentials represents the Git credentials.
type GitCredentials struct {
	Provider string `yaml:"provider"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

// Project represents a deployment project.
type Project struct {
	Name              string `yaml:"name"`
	Repository        string `yaml:"repository"`
	Branch            string `yaml:"branch"`
	DockerComposePath string `yaml:"docker_compose_path"`
}

// Config represents the overall YAML configuration.
type Config struct {
	Environment    Environment    `yaml:"environment"`
	GitCredentials GitCredentials `yaml:"git_credentials"`
	Projects       []Project      `yaml:"projects"`
}

func Init(filePath string) (Config, error) {

	yamlContent, err := ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse YAML configuration
	parsedConfig, err := ParseConfig(yamlContent)
	if err != nil {
		log.Fatal(err)
	}

	return parsedConfig, nil
	//return true, errors.New("No file found on path " + file)
}

// ParseConfig parses the YAML configuration file.
func ParseConfig(yamlData []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
