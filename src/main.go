package main

import (
	"cli-todos/cmd"
	"cli-todos/pkg"
)

func main() {
	pkg.Initialize()
	cmd.Execute()
}
