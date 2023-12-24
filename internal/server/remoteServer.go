package server

type AuthMethod int

const (
	// Password authentication method
	Password AuthMethod = iota

	// PublicKey authentication method
	PublicKey
)

// RemoteServer represents the details of the remote server.
type RemoteServer struct {
	Host       string
	Port       string
	Username   string
	Key        string
	AuthMethod AuthMethod
}
