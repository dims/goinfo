package main

import (
	"debug/buildinfo"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "goinfo <file>", Run: run}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	files := cmd.Flags().Args()

	if len(files) == 0 {
		fmt.Println("please specify a file")
		os.Exit(1)
	}

	result, err := buildinfo.ReadFile(files[0])
	if err != nil {
		fmt.Println("Error parsing file:", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
