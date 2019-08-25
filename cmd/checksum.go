package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"

	"github.com/spf13/cobra"
)

var checkfilename string
var md5Algorithm bool
var sha1Algorithm bool
var sha256Algorithm bool

//Add a subcommand checksum under fopd
func init() {
	RootCmd.AddCommand(checksumCmd)

	//flag --file -f
	checksumCmd.Flags().StringVarP(&checkfilename, "file", "f", "", "the input file")
	checksumCmd.MarkFlagRequired("file")

	//flag --md5, --sha1, --sha256
	checksumCmd.Flags().BoolVar(&md5Algorithm,"md5",false, "print the checksum of the file with md5 algorithm")
	checksumCmd.Flags().BoolVar(&sha1Algorithm,"sha1",false, "print the checksum of the file with sha1 algorithm")
	checksumCmd.Flags().BoolVar(&sha256Algorithm,"sha256",false, "print the checksum of the file with sha256 algorithm")

}

//Implement checksum and its function
var checksumCmd = &cobra.Command{
	Use:   "checksum",
	Short: "Print checksum of file",
	RunE: func(cmd *cobra.Command, args []string) error {

		//read the file
		data, err := ioutil.ReadFile(checkfilename)
		if err != nil {
			fmt.Printf("error: '%s'%s\n", checkfilename, strings.Split(err.Error(), ":")[1])
			return nil
		}

		//distinguish different flag for each algorithm, and print the result of the file's checksum
		if md5Algorithm == true && sha1Algorithm == false && sha256Algorithm == false {

			fmt.Printf("%x\n", md5.Sum(data))

		} else if md5Algorithm == false && sha1Algorithm == true && sha256Algorithm == false {

			fmt.Printf("%x\n", sha1.Sum(data))

		} else if md5Algorithm == false && sha1Algorithm == false && sha256Algorithm == true {

			fmt.Printf("%x\n", sha256.Sum256(data))

		} else {

			fmt.Println("error: use one and only one of --md5, --sha1 and --sha256 in checksum")

		}

		return nil
	},
}
