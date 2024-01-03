package server

import "github.com/jinzhu/gorm"

type AuthMethod int

const (
	// Password authentication method
	Password AuthMethod = iota

	// PublicKey authentication method
	PublicKey
)

// RemoteServer represents the details of the remote server.
type RemoteServer struct {
	gorm.Model
	Name       string
	Host       string
	Port       string
	Username   string
	Key        string
	AuthMethod AuthMethod
}
