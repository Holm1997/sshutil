package utils

import (
	"flag"
	"os"
)

var FilePath = flag.String("f", "", "указывает путь до конфигурационного .YAML файла ")

var Cmd = flag.String("c", "", "передает shell-команды на удаленный хост")

var Src = flag.String("src", "", "передает файл который необходимо скопировать на локальном хосте")

var Dst = flag.String("dst", "", "указывает файл в который необходимо скопировать на удаленном хосте")

var CmdExists = false

var SrcDstExists = false

func ParseFlags() {
	flag.Parse()

	fl := flag.NFlag()

	if fl == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	} else if *Cmd != "" {
		CmdExists = true
	}

	if *Src != "" {
		if *Dst != "" {
			SrcDstExists = true
		}
	} else if *Dst != "" {
		if *Src != "" {
			SrcDstExists = true
		}
	}
}
