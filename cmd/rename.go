package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename [oldPath] [newPath]",
	Short: "rename a file or directory",
	Long:  `rename a file or directory`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]
		renameFileOrDir(oldPath, newPath)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func renameFileOrDir(oldPath, newPath string) {
	if err := os.Rename(oldPath, newPath); err != nil {
		fmt.Println("error renaming:", err)
		return
	}

	fmt.Println("Renamed", oldPath, "to", newPath)
}
