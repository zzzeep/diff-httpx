/*
Copyright © 2023 zzzeep
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zzzeep/diff-httpx/diff"
	"github.com/zzzeep/diff-httpx/output"
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

		changes := diff.GetChanges(oldRecords, newRecords)
		output.PrintTable(changes)
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
	rootCmd.Flags().BoolVar(&output.Options.NoColor, "nc", false, "no color")
	rootCmd.Flags().BoolVar(&output.Options.NoColor, "nt", false, "don't shorten long fields")
	rootCmd.Flags().UintVar(&output.Options.FilterCode, "fs", 0, "filter by status code")

	rootCmd.Flags().BoolVarP(&output.Options.IPs, "ip", "i", false, "show ip changes")
	rootCmd.Flags().BoolVarP(&output.Options.Port, "port", "p", false, "show port changes")
	rootCmd.Flags().BoolVarP(&output.Options.Webserver, "web-server", "w", false, "show web-server changes")
	rootCmd.Flags().BoolVarP(&output.Options.StatusCode, "status-code", "s", false, "show status code changes")
	rootCmd.Flags().BoolVarP(&output.Options.Title, "title", "t", false, "show title changes")
	rootCmd.Flags().BoolVarP(&output.Options.ContentType, "content-type", "c", false, "show content type changes")
	rootCmd.Flags().BoolVarP(&output.Options.ContentLength, "content-length", "l", false, "show content length changes")
	rootCmd.Flags().BoolVarP(&output.Options.Hash, "hash", "x", false, "show body or header hash changes")
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
