package main

import (
	"flag"
	etcd2 "flagVars/flaglookup/etcd"
)

/*
	不需要我们显示注入“flag vars”了，我们只需按照flag提供的方法在其他package中读取对应flag变量的值即可

*/

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

func main() {
	flag.Parse()

	go etcd2.Proxy()
}
