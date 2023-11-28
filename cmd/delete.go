package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a file with undo support",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		deleteFile(filePath)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type DeletedFileInfo struct {
	Path       string
	DeleteTime time.Time
}

var deletedFiles []DeletedFileInfo

func deleteFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		fmt.Println("Error deleting file", err)
		return
	}

	fmt.Printf("File '%v' deleted successfully", filepath)

	//log deleted file
	deletedFiles = append(deletedFiles, DeletedFileInfo{
		Path:       filepath,
		DeleteTime: time.Now(),
	})
}
