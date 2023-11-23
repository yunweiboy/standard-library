package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	endpoints string
	username  string
	password  string
)

func init() {
	// 自定义string参数变量
	flag.StringVar(&endpoints, "endpoints", "127.0.0.1:2379", "etcd endpoints")
	flag.StringVar(&username, "username", "", "etcd client user")
	flag.StringVar(&password, "password", "", "etcd client user password")
}

// 定义usage
func usage() {
	fmt.Println("flagdemo-app is daemon application which provides xxx service.")
	fmt.Println("Usage of flagdemo-app:")
	fmt.Println("\t flagdemo-app [options]:")
	fmt.Println("The options are:")

	flag.PrintDefaults() // 按格式 输出command flag信息
}

func main() {
	flag.Usage = usage
	flag.Parse()

	fmt.Println(endpoints, username, password)
	fmt.Println(os.Args, len(os.Args)) // [./main -endpoints 127.0.0.1 -username kxc -password 123] 7
}
