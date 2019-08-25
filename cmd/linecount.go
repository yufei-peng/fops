package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var filename string

//Add a subcommand linecount under fopd
func init() {
	RootCmd.AddCommand(linecountCmd)

	//flag --file -f
	linecountCmd.Flags().StringVarP(&filename, "file", "f", "", "the input file")
	linecountCmd.MarkFlagRequired("file")
}

//Implement linecount and its function
var linecountCmd = &cobra.Command{
	Use:   "linecount",
	Short: "Print line count of file",
	RunE: func(cmd *cobra.Command, args []string) error {

		//check if the argument of file flag is legal: not a file, exist, not a binary file.
		fileinfo, err := os.Stat(filename)
		if err != nil {
			fmt.Printf("error:%s\n", strings.Split(err.Error(), ":")[1])
			return nil
		}
		if fileinfo.IsDir() == true {
			fmt.Printf("error: Expected file got directory '%s'\n", filename)
			return nil
		} else if strings.Index(fileinfo.Mode().String(), "x") != -1 {
			fmt.Printf("error: Cannot do linecount for binary file '%s'\n", filename)
			return nil
		}

		//read the file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("error: ", err)
			return nil
		}

		//count how many lines does this file have
		fmt.Println(len(strings.Split(string(data), "\n")))
		return nil

	},
}
