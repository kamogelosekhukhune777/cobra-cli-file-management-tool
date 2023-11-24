/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// copyCmd represents the copy command
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[0]
		destFile := args[1]
		copyFile(sourceFile, destFile)
	},
}

func init() {
	rootCmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func copyFile(sourcePath, destPath string) {
	input, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("Error reading source file:", err)
		return
	}

	err = os.WriteFile(destPath, input, 0644)
	if err != nil {
		fmt.Println("Error copying fing file", err)
		return
	}

	fmt.Printf("File %v copied to %v successfully", sourcePath, destPath)
}
