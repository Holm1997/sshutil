package utils

import (
	"flag"
	"os"
)

var FilePath = flag.String("f", "", "указывает путь до конфигурационного .YAML файла ")

var Cmd = flag.String("c", "", "передает shell-команды на удаленный хост")

func ParseFlags() {
	flag.Parse()

	fl := flag.NFlag()

	if fl == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
