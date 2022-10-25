package main

import (
	"fmt"
	"hsec/common"
)

func init() {
	// Before main
	// Intruduce
	fmt.Printf("This is a host sec check tool.\nVersion: %s\nAuthor: %s\n", common.Version, common.Author)
}

func main() {
	common.GetProcess()

}
