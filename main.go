package main

import (
	"scf_scanner_client/modules"
)

func main() {
	S := new(modules.Scanner)
	S.Run()
}
