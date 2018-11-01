package main

import (
	"fmt"

	VirtualCloudCore "github.com/virtual-cloud/core"
)

func main() {
	var man VirtualCloudCore.Man
	fmt.Println(man)
}

func commandLs() {
	VirtualCloudCore.CommandLs()
}
