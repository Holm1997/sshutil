package utils

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func ExecuteCmd(cmd string, hostname string, config *ssh.ClientConfig) {
	client, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(cmd); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
