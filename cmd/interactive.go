package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Run in interactive mode.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		interactiveMode()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interactiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interactiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to interactive Mode! type 'help' for available commands.")

	for {
		fmt.Println(">")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		switch command {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("- help: Show available commands.")
			fmt.Println("- fileinfo [path]: Display file or directory information.")
			fmt.Println("- zip [source] [destination]: create a zip archieve")
			fmt.Println("- unzip [archieve] [destination]: Extract a zip archieve.")
			//other cases...
		case "fileinfo":
			if len(args) != 3 {
				fmt.Println("Usage: fileinfo[path]")
				continue
			}
			fileInfo(args[1])
		case "zip":
			if len(args) != 3 {
				fmt.Println("Usage: zip[source] [destination]")
				continue
			}
			createZipArchieve(args[1], args[2])
		case "unzip":
			if len(args) != 3 {
				fmt.Println("Usage: unzip [archive] [destination]")
				continue
			}
			extractZipArchieve(args[1], args[2])
			//other cases.......
		case "exit":
			fmt.Println("Exiting interating Mode.")
			return
		default:
			fmt.Println("Invalid command, Type 'help' for available commads.")
		}
	}
}
