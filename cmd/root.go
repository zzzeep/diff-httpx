/*
Copyright © 2023 zzzeep
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zzzeep/diff-httpx/parser"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "diff-httpx",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("provide two files to generate a diff report")
		}

		oldRecords, err := readHttpxJson(args[0])
		if err != nil {
			return
		}

		newRecords, err := readHttpxJson(args[1])
		if err != nil {
			return
		}

		fmt.Println(oldRecords[0])
		fmt.Println(newRecords[0])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.diff-httpx.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	var b bool
	rootCmd.Flags().BoolVar(&b, "test", false, "")
}

func readHttpxJson(path string) ([]parser.HttpxRecord, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("error reading file:", path)
		return nil, err
	}
	content := string(bytes)
	records := parser.ParseHttpx(content)
	return records, nil
}
