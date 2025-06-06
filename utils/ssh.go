package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func ExecuteCmd(cmd string, hostname string, port string, config *ssh.ClientConfig) error {
	client, err := ssh.Dial("tcp", hostname+":"+port, config)
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
	return nil
}

func CopyFile(src string, dst string, hostname string, port string, config *ssh.ClientConfig) error {
	client, _ := ssh.Dial("tcp", hostname+":22", config)

	defer client.Close()

	// open an SFTP session over an existing ssh connection.
	sftp, err := sftp.NewClient(client)
	if err != nil {
		return err
	}
	defer sftp.Close()

	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := sftp.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// write to file
	if _, err := dstFile.ReadFrom(srcFile); err != nil {
		return err
	}

	fmt.Printf("Файл %v успешно передан\n", srcFile.Name())
	return nil
}
