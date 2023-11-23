package main

import (
	"flag"
	etcd2 "flagVars/parampass/etcd"
	"time"
)

/*
	将Parse后的flag变量以参数的形式、以某种init的方式传给其他要使用这些变量的包。

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

	go etcd2.Etcd(endpoints, username, password)

	time.Sleep(5 * time.Second)
}
