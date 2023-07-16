package main

import (
	"github.com/falo2/ma/cmd" // import the cmd package
)

func main() {
	cmd.Init()    // initialize the commands and flags
	cmd.Execute() // execute the commands
}
