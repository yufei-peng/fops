package cmd

import (
	"fmt"
	"log"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

//Add a subcommand version under fopd
func init() {
	RootCmd.AddCommand(versionCmd)
}

//Setup subcommand version. Read from version file, and show the current version of fops
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version info",
	Run: func(cmd *cobra.Command, args []string) {

		box := packr.NewBox("../../fops")
		data, err := box.FindString("version")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("fops", data)
	},

}
