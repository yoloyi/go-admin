package main

import (
	"fmt"
	"monitor/configs"
)

func main() {
	configs.InitConfig()
	fmt.Println(configs.GetDbConfig())
}
