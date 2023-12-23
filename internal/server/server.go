package server

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
)

// CopyFileToRemote copies a local file to a remote server using SCP.
func (server RemoteServer) CopyFileToRemote(localFilePath, remoteFilePath string) error {
	// Read the private key file
	key, err := ioutil.ReadFile(server.KeyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %v", err)
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Set up the SSH client configuration
	config := &ssh.ClientConfig{
		User: server.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Change in production for security
	}

	// Connect to the remote server over SSH
	sshClient, err := ssh.Dial("tcp", server.Host+":"+server.Port, config)
	if err != nil {
		return fmt.Errorf("failed to connect to SSH server: %v", err)
	}
	defer sshClient.Close()

	// Open an SFTP session on top of the existing SSH connection
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return fmt.Errorf("failed to create SFTP client: %v", err)
	}
	defer sftpClient.Close()

	// Open the local file for reading
	localFile, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to open local file: %v", err)
	}
	defer localFile.Close()

	// Create the remote file for writing
	remoteFile, err := sftpClient.Create(remoteFilePath)
	if err != nil {
		return fmt.Errorf("failed to create remote file: %v", err)
	}
	defer remoteFile.Close()

	// Copy the contents of the local file to the remote file
	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		return fmt.Errorf("failed to copy file contents: %v", err)
	}

	fmt.Printf("File '%s' copied to remote server '%s:%s'\n", localFilePath, server.Host, remoteFilePath)

	return nil
}

func (server RemoteServer) UnzipFileOnRemote(remoteFilePath, destinationDir string) error {
	// Read the private key file
	key, err := ioutil.ReadFile(server.KeyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %v", err)
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Set up the SSH client configuration
	config := &ssh.ClientConfig{
		User: server.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Change in production for security
	}

	// Connect to the remote server over SSH
	sshClient, err := ssh.Dial("tcp", server.Host+":"+server.Port, config)
	if err != nil {
		return fmt.Errorf("failed to connect to SSH server: %v", err)
	}
	defer sshClient.Close()

	// Create a session
	session, err := sshClient.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Build the unzip command
	cmd := fmt.Sprintf("unzip -o %s -d %s", remoteFilePath, destinationDir)

	// Execute the command on the remote server
	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to run unzip command: %v", err)
	}

	fmt.Printf("File '%s' unzipped to remote server '%s:%s'\n", remoteFilePath, server.Host, destinationDir)

	return nil
}
