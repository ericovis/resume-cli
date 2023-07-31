/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// liveCmd represents the live command
var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "Opens a live preview of your resume in your browser",
	Long:  `Opens a live preview of your resume in your browser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("live called")
	},
}

func init() {
	rootCmd.AddCommand(liveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// liveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// liveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
