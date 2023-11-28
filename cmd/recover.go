package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// recoverCmd represents the recover command
var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "recover (undelete) a file",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		recoverFile(filePath)
		fmt.Println("recover called")
	},
}

func init() {
	rootCmd.AddCommand(recoverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recoverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recoverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func recoverFile(filePath string) {
	for i, fileInfo := range deletedFiles {
		if fileInfo.Path == filePath {
			//restore file

			err := os.Rename(filePath+".deleted", filePath)
			if err != nil {
				fmt.Println("error recovering file:", err)
				return
			}

			//remove entry from log
			deletedFiles = append(deletedFiles[:i], deletedFiles[i+1:]...)
			fmt.Println("File recovered:", filePath)

		}
	}
	fmt.Println("File not found in deletion history.")
}
