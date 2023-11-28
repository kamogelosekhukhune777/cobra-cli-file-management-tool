package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip [source] [destination]",
	Short: "Create a zip archive",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sourcePath := args[0]
		destPath := args[1]
		createZipArchieve(sourcePath, destPath)
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createZipArchieve(sourcePath, destPath string) {
	zipFile, err := os.Create(destPath)
	if err != nil {
		fmt.Println("Error creating zip file:", err)
		return
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(sourcePath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name, _ = filepath.Rel(sourcePath, filePath)
		if info.IsDir() {
			header.Name += "/"
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			_, err := io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	fmt.Println("Zip archive created:", destPath)
}
