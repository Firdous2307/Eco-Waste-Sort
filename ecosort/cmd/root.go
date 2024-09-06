package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ecosort",
    Short: "EcoSort assists you in sorting waste correctly.",
    Long:  `EcoSort is a CLI tool that helps you sort waste into recyclable, compostable, and non-recyclable categories.`,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        // Handle error appropriately
    }
}