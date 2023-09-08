package main

import (
	"fmt"

	"openeuler.org/PilotGo/atune-plugin/utils"
)

func main() {
	info := utils.GetTuneInfo("gcc_compile")
	fmt.Printf("%#v", info)
}
