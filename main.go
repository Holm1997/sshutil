package main

import (
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
		config := utils.CreateClientConfig(host.Name, host.Pass)
		utils.ExecuteCmd(*cmd, host.Name, config)
		err := utils.CopyFile(*src, *dst, host.Name, config)
		if err != nil {
			log.Fatal(err)
		}
	}
}
