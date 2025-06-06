package main

import (
	"fmt"
	"log"

	"github.com/Holm1997/sshutil/utils"
)

func main() {
	utils.ParseFlags()
	cmd := utils.Cmd
	file := utils.FilePath
	data := utils.ParseYamlFile(file)
	src := utils.Src
	dst := utils.Dst

	for _, host := range data.Hosts {
		config := utils.CreateClientConfig(host.Login, host.Pass)
		fmt.Printf("-------[ %v ]-------\n", host.Host)

		if host.Port == "" {
			host.Port = "22"
		}

		if utils.CmdExists {
			utils.ExecuteCmd(*cmd, host.Host, host.Port, config)
		}

		if utils.SrcDstExists {
			err := utils.CopyFile(*src, *dst, host.Host, host.Port, config)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("-----------------------")
	}
}
