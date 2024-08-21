/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CobraQRCodeGenerator",
	Short: "A CLI tool for generating QR codes",
	Long: `CobraQRCodeGenerator is a command-line tool designed to generate QR codes.

This application uses the Cobra CLI library to provide a user-friendly interface for generating QR codes from provided text content. 

Features:
- Generate QR codes in various formats.
- Customize output file name and QR code size.
- Easily integrate into your projects or use for quick QR code generation.

With CobraQRCodeGenerator, you can quickly create QR codes for URLs, text, or other data directly from your terminal or command line interface.

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CobraQRCodeGenerator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
