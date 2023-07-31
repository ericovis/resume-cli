/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	r "github.com/ericovis/resume-cli/src/resume"
	"github.com/spf13/cobra"
)

var resumeFile string
var resume r.Resume

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "resume-cli",
	Short: "A CLI tool for building a resumés from JSON and YAML files.",
	Long: `A CLI tool for building a resumés from JSON and YAML files.
	Based on the JSON Resume Schema (https://jsonresume.org/schema/)`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		if resumeFile == "" {
			cwd, _ := os.Getwd()
			resumeFile, err = r.FindResumeFileOnDir(cwd)

			ExitOnError(err)
		}

		err = resume.Load(resumeFile)
		ExitOnError(err)
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	ExitOnError(err)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&resumeFile, "file", "f", "", "path for the resume file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
