package fops

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//Define the command
var RootCmd = &cobra.Command{
	Use: "fops",
	Short: "File Ops",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
