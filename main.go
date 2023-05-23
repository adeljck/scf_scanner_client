package main

import (
	"fmt"
	"scf_scanner_client/modules"
)

func main() {
	modules.S.Run()
	fmt.Println(modules.S.Results)
}
