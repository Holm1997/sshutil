package main

import (
	"github.com/Holm1997/sshutil/utils"
	"golang.org/x/crypto/ssh"
)

func main() {
	utils.ParseFlags()
	cmd := utils.Cmd
	file := utils.FilePath
	data := utils.ParseYamlFile(file)

	for _, host := range data.Hosts {
		config := &ssh.ClientConfig{
			User: host.Login,
			Auth: []ssh.AuthMethod{
				ssh.Password(host.Pass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
		utils.ExecuteCmd(*cmd, host.Name, config)
	}
}
