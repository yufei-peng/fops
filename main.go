package main

import (
	"fmt"
	"os"

	fops "github.com/yufei-peng/fops/cmd"
)

//main.go only use for initializing Cobra
func main() {
	if err := fops.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}