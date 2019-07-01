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
	Long:  "Returns a Hash of the given file",
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
		filepath := path + "/" + filename

		if _, err := os.Stat(filepath); err == nil {
			hashProcessor := *hashprocess.GetHashFunction(filepath, strings.ToLower(function))

			if hashProcessor != nil {
				if hash, err := hashProcessor.Process(); err == nil {
					fmt.Println(filename + ": " + hash)
				} else {
					fmt.Println("Error Hashing File: " + err.Error())
				}
			} else {
				fmt.Println("Hash Function Not Supported")
			}
		}
	}
}
