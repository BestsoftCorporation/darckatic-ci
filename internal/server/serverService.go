package server

import (
	"darkatic-ci/internal/db"
	"fmt"
)

func init() {
	db.DB.AutoMigrate(&RemoteServer{})
}

func GetServerByHostname(hostname string) (*RemoteServer, error) {
	var server RemoteServer
	if err := db.DB.Where("host = ?", hostname).First(&server).Error; err != nil {
		return nil, fmt.Errorf("server not found")
	}

	return &server, nil
}

func GetServers() *[]RemoteServer {
	var servers []RemoteServer
	db.DB.Find(&servers)
	return &servers
}
