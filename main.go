package main

import (
	"cloud-task/core"
	"cloud-task/global"
	"fmt"
)

func main() {
	core.Init()
	fmt.Println(global.GVA_CONFIG.System.Version)
}
