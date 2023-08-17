package main

import (
	"github.com/jackaitken/toto/cmd"
	"github.com/jackaitken/toto/pkg"
)

func main() {
	pkg.Initialize()
	cmd.Execute()
}
