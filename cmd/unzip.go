package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip [archieve]",
	Short: "Extract a zip archieve",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		archivePath := args[0]
		destPath := args[1]
		extractZipArchieve(archivePath, destPath)
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unzipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func extractZipArchieve(archivePath, destPath string) {
	zipFile, err := zip.OpenReader(archivePath)
	if err != nil {
		fmt.Println("error opening zip:", err)
		return
	}
	defer zipFile.Close()

	for _, file := range zipFile.File {
		filePath := filepath.Join(destPath, file.Name)

		if file.FileInfo().IsDir() {
			os.Mkdir(filePath, os.ModePerm)
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			fmt.Println("Error opening file in archive:", err)
		}
		defer fileReader.Close()

		outFile, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating extracted file:", err)
			return
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, fileReader)
		if err != nil {
			fmt.Println("Error extracting file:", err)
			return
		}
	}

	fmt.Println("Zip archive extracted to:", destPath)
}
