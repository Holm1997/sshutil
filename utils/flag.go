package utils

import (
	"flag"
	"os"
)

var FilePath = flag.String("f", "", "указывает путь до конфигурационного .YAML файла ")

var Cmd = flag.String("c", "", "передает shell-команды на удаленный хост")

var Src = flag.String("src", "", "передает файл который необходимо скопировать на локальном хосте")

var Dst = flag.String("dst", "", "указывает файл в который необходимо скопировать на удаленном хосте")

func ParseFlags() {
	flag.Parse()

	fl := flag.NFlag()

	if fl == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}
