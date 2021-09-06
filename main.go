package main

import (
	"fmt"

	"github.com/clh021/crud-api/cmd"
)

var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)
	cmd.Execute()
}
