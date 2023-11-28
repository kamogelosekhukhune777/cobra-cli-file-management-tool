package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// bulkCmd represents the bulk command
var bulkCmd = &cobra.Command{
	Use:   "bulk [opertion] [paths...]",
	Short: "Perform bulk operations on files or directories",
	Long:  ``,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		operation := args[0]

		//path arguments(file or directories)
		paths := args[1 : len(args)-1]

		//destination path (if required for copy or move operations)
		destPath := args[len(args)-1]

		bulkOperation(operation, paths, destPath)
	},
}

func init() {
	rootCmd.AddCommand(bulkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bulkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bulkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func bulkOperation(operation string, paths []string, destpath string) {
	switch operation {
	case "copy":
		for _, path := range paths {
			copyFile(path, destpath)
		}

	case "move":
		for _, path := range paths {
			moveFileWithUndo(path, destpath)
		}
	case "delete":
		for _, path := range paths {
			deleteFile(path)
		}
		fmt.Println("Bulk delete operation completed")
	default:
		fmt.Println("inalid operation")
	}
}
