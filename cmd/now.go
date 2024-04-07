/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takurooo/aatime/internal/aa"
)

const (
	nowShortDescription = "Display current time in ASCII art."
	nowLongDescription  = `Display current time in ASCII art.`
)

func newCmdNow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "now",
		Short: nowShortDescription,
		Long:  nowLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNow(cmd, args)
		},
	}
	cmd.Flags().StringP("font", "f", "invita", "ASCII art font. (invita or train or bulbhead)")
	cmd.Flags().BoolP("utc", "u", false, "Display current time in UTC")
	return cmd
}

func runNow(cmd *cobra.Command, args []string) error {
	font, err := cmd.Flags().GetString("font")
	if err != nil {
		return err
	}
	utc, err := cmd.Flags().GetBool("utc")
	if err != nil {
		return err
	}
	fmt.Println(aa.Now(aa.FontFromString(font), utc))
	return nil
}
