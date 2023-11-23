package main

import (
	"flag"
	"flagVars/globalVars/config"
	"flagVars/globalVars/etcd"
	"fmt"
	"os"
	"time"
)

/*
	要注意这些全局变量值在Go包初始化过程的顺序，比如：如果在etcd包的init函数中使用这些全局变量，那么你得到的各个变量值将为空值，
因为etcd包的init函数在main.init和main.main之前执行，这个时候绑定和Parse都还未执行。
*/

func init() {
	// 自定义string参数变量
	flag.StringVar(&config.Endpoints, "endpoints", "127.0.0.1:2379", "etcd endpoints")
	flag.StringVar(&config.Username, "username", "", "etcd client user")
	flag.StringVar(&config.Password, "password", "", "etcd client user password")
}

func main() {
	flag.Parse()
	fmt.Println(os.Args[0])                    // ./main
	fmt.Println(os.Args[1:], len(os.Args[1:])) // [-endpoints 127.0.0.1 -username kxc -password 123] 6

	go etcd.Etcd()

	time.Sleep(5 * time.Second)
}
