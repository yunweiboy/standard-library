package etcd

import (
	"flagVars/globalVars/config"
	"fmt"
)

func Etcd() {
	fmt.Println(config.Endpoints, config.Username, config.Password)
}
