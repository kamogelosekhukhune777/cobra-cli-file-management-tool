package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move [oldPath] [newPath]",
	Short: "move (rename) a file or directory with undo support",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]
		moveFileWithUndo(oldPath, newPath)
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type RenamedFileInfo struct {
	OldPath    string
	NewPath    string
	RenameTime time.Time
}

var renamedFiles []RenamedFileInfo

func moveFileWithUndo(oldPath, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("error moving file:", err)
		return
	}
	fmt.Println("file moved:", oldPath, "->", newPath)

	renamedFiles = append(renamedFiles, RenamedFileInfo{
		OldPath:    oldPath,
		NewPath:    newPath,
		RenameTime: time.Now(),
	})
}
