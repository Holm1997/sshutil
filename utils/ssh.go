package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"

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

func CopyFile(src string, dst string, hostname string, config *ssh.ClientConfig) error {
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
	// Отправка файла
	sourceFile, err := os.ReadFile(src)
	if err != nil {
		panic(err)
	}

	destinationPath := dst

	// Открытие канала для передачи данных
	_, err = session.Output(fmt.Sprintf("echo %s > %s", sourceFile, destinationPath))
	if err != nil {
		panic(err)
	}

	fmt.Println("Файл успешно передан")
	return nil
}
