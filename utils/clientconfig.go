package utils

import "golang.org/x/crypto/ssh"

func CreateClientConfig(host string, pass string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: host,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}
