package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/MichaelWittgreffe/FileHash/hashprocess"
	"github.com/spf13/cobra"
)

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

func rootCmdHandler(cmd *cobra.Command, args []string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error Getting Working Directory")
		return
	}

	if len(args) == 0 {
		fmt.Println("Filename Param Not Given")
		return
	}

	filename := args[0]
	filepath := path + "/" + filename
	var function string

	if len(args) == 2 {
		function = args[1]
	} else {
		function = "md5"
	}

	_, err = os.Stat(filepath)
	if err != nil {
		fmt.Println("File Not Found")
		return
	}

	hashProcessor := hashprocess.GetHashProcessor(filepath, strings.ToLower(function))
	if hashProcessor == nil {
		fmt.Println("Hash Function Not Supported")
		return
	}

	hash, err := hashProcessor.Process()
	if err != nil {
		fmt.Printf("Error Hashing File: %s\n", err.Error())
	}

	fmt.Println(filename + ": " + hash)
}
