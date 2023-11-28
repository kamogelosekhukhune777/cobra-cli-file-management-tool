package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// chmodCmd represents the chmod command
var chmodCmd = &cobra.Command{
	Use:   "chmod [file] [permissions]",
	Short: "Change file permissions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		permissions := args[1]

		modifyPermissions(filePath, permissions)
	},
}

func init() {
	rootCmd.AddCommand(chmodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chmodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chmodCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func modifyPermissions(filePath, permString string) {
	perm, err := strconv.ParseUint(permString, 8, 32)
	if err != nil {
		fmt.Println("Permission value:", err)
		return
	}

	if err := os.Chmod(filePath, os.FileMode(perm)); err != nil {
		fmt.Println("Error changing permissions:", err)
		return
	}

	fmt.Println("Permisssions Changed successfully")
}
