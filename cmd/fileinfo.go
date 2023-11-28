package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

// fileinfoCmd represents the fileinfo command
var fileinfoCmd = &cobra.Command{
	Use:   "fileinfo [path]",
	Short: "display file or directory information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		fileInfo(path)
	},
}

func init() {
	rootCmd.AddCommand(fileinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fileinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fileinfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func fileInfo(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Println("File/Directory:", path)
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permission:", fileInfo.Mode().String())
	fmt.Println("Last Modified:", fileInfo.ModTime().Format(time.RFC822))
	fmt.Println("Is Directory:", fileInfo.IsDir())

	if fileInfo.IsDir() {
		//list content of the directory
		fmt.Println("\nContents")
		files, err := filepath.Glob(filepath.Join(path, "*"))
		if err != nil {
			fmt.Println("error listing directory contents", err)
			return
		}

		for _, file := range files {
			fmt.Println(file)
		}

	}
}
