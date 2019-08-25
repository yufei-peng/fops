package main

import (
	"fops/cmd"
	"fmt"
	"os"
)

//main.go only use for initializing Cobra
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}