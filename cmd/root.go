/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	rootShortDescription = "ASCII art time command."
	rootLongDescription  = `ASCII art time command.`
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aatime",
	Short: rootShortDescription,
	Long:  rootLongDescription,
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
	rootCmd.AddCommand(newCmdNow())
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
