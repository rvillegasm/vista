package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "vista",
    Short: "Vista CLI - context engineering assistant",
    Long:  "Vista helps generate context engineering files for AI coding agents.",
}

// Execute runs the root command and exits with an error code on failure.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func init() {
    rootCmd.AddCommand(setupCmd)
    rootCmd.AddCommand(newCmd)
} 