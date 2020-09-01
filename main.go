package main

import (
	"flag"
	"log"
)

var name string

func main() {
	// 获取指令参数，os.args[1:]
	flag.Parse()
	// 创建一个go新的子命令集
	goCMd := flag.NewFlagSet("go", flag.ExitOnError)
	goCMd.StringVar(&name, "name", "go 语言", "帮助信息")
	// 创建一个php新的子命令集
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "php语言", "帮助信息")
	// 获取指令参数
	args := flag.Args()
	switch args[0] {
	//如果是go指令，走gocmd
	case "go":
		_ = goCMd.Parse(args[1:])
	case "php":
		//如果是php指令，走phpcmd
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}
