package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/MichaelWittgreffe/FileHash/hashprocess"
	"github.com/spf13/cobra"
)

//---------------------------------------
// Define CLI
//---------------------------------------

var rootCmd = &cobra.Command{
	Use:   "app [filename] [hash function]",
	Short: "Hash A File",
	Long:  "Returns a Hash of the given file in the format \"filename: hashstring\"",
	Args:  cobra.MinimumNArgs(1),
	Run:   rootCmdHandler,
}

//Execute runs the cli app
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//---------------------------------------
// Define Handler Functions
//---------------------------------------

func rootCmdHandler(cmd *cobra.Command, args []string) {
	if path, err := os.Getwd(); err == nil {
		filename := args[0]
		function := args[1]

		if len(filename) > 0 {
			filepath := path + "/" + filename

			//default hash func is MD5
			if len(function) == 0 {
				function = "md5"
			}

			if _, err := os.Stat(filepath); err == nil {
				hashProcessor := hashprocess.GetHashProcessor(filepath, strings.ToLower(function))

				if hashProcessor != nil {
					if hash, err := hashProcessor.Process(); err == nil {
						fmt.Println(filename + ": " + hash)
					} else {
						fmt.Println("Error Hashing File: " + err.Error())
					}
				} else {
					fmt.Println("Hash Function Not Supported")
				}
			} else {
				fmt.Println("File Not Found")
			}
		} else {
			fmt.Println("Filename Param Not Given")
		}
	}
}
