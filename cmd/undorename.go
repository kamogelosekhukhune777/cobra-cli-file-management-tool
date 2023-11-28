package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// undorenameCmd represents the undorename command
var undorenameCmd = &cobra.Command{
	Use:   "undorename [path]",
	Short: "undo a file or directory name",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newPath := args[0]
		undoRenameFile(newPath)
	},
}

func init() {
	rootCmd.AddCommand(undorenameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undorenameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undorenameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func undoRenameFile(newPath string) {
	for i := len(renamedFiles) - 1; i >= 0; i-- {
		if renamedFiles[i].NewPath == newPath {
			//revert the renaming action
			err := os.Rename(newPath, renamedFiles[i].OldPath)
			if err != nil {
				fmt.Println("error undoing rename:", err)
				return
			}
			//remove the entry
			renamedFiles = append(renamedFiles[:i], renamedFiles[i+1:]...)
			fmt.Println("Rename undone:", newPath, "->", renamedFiles[i].OldPath)
		}
	}
	fmt.Println("Rename not found in history")
}
