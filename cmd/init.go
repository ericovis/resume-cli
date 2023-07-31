/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var output string
var json bool = false

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:              "init",
	Short:            "Creates a resume data file on the current (or specified) directory.",
	Long:             `Creates a resume data file on the current (or specified) directory.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		resume.Init(output, json)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	cwd, _ := os.Getwd()
	initCmd.Flags().StringVarP(&output, "output", "o", cwd, "specifies the output directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolVarP(&json, "json", "j", json, "Uses json for resume data")
}
