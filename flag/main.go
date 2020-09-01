package main

import (
	"flag"
	"log"
	"os"
)

var (
	name string
	age  int
)

func init() {
	log.Println("func init running...")
	flag.StringVar(&name, "name", "张三", "输入名字")
	flag.IntVar(&age, "age", 30, "输入年龄")
}

func main() {
	flag.Parse()
	args := flag.Args()
	log.Printf("%s今年%d岁了", name, age)
	log.Println(args)
	log.Println(os.Args)
	log.Println("func main running...")
}
