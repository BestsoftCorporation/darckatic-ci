package server

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"

	"github.com/docker/docker/client"
)

// DockerCompose represents the structure of a Docker Compose file.
type DockerCompose struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
}

// Service represents the structure of a Docker Compose service.
type Service struct {
	Environment map[string]string `yaml:"environment"`
}

// LogDockerVersion logs the Docker version on a remote server using the Docker client library.
func (server RemoteServer) LogDockerVersion() error {
	// Create a Docker client using SSH
	cli, err := client.NewClientWithOpts(
		client.WithHost(fmt.Sprintf("ssh://%s:%s", server.Host, server.Port)),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return fmt.Errorf("failed to create Docker client: %v", err)
	}

	// Get the Docker version
	version, err := cli.ServerVersion(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get Docker version: %v", err)
	}

	// Log the Docker version
	fmt.Printf("Docker Version on %s:\n", server.Host)
	fmt.Printf("  Version: %s\n", version.Version)
	fmt.Printf("  API Version: %s\n", version.APIVersion)

	return nil
}

// ListDockerComposeEnvVars parses Docker Compose environment variables and lists them.
func ListDockerComposeEnvVars(filePath string) error {
	// Read the Docker Compose file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read Docker Compose file: %v", err)
	}

	// Parse the YAML content
	var dockerCompose DockerCompose
	err = yaml.Unmarshal(content, &dockerCompose)
	if err != nil {
		return fmt.Errorf("failed to parse Docker Compose YAML: %v", err)
	}

	fmt.Println("Docker Compose Environment Variables:")
	for _, service := range dockerCompose.Services {
		for key, value := range service.Environment {
			fmt.Printf("%s=%s\n", key, value)
		}
	}

	return nil
}
