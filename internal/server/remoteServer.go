package server

// RemoteServer represents the details of the remote server.
type RemoteServer struct {
	Host     string
	Port     string
	Username string
	KeyPath  string // Path to the private key file for SSH authentication
}
