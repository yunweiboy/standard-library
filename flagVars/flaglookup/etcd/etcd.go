package etcd

import (
	"flag"
	"fmt"
)

func Proxy() {
	endpoints := flag.Lookup("endpoints").Value.(flag.Getter).Get().(string)
	username := flag.Lookup("username").Value.(flag.Getter).Get().(string)
	password := flag.Lookup("password").Value.(flag.Getter).Get().(string)

	fmt.Println(endpoints, username, password)
}
