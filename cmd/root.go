package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "vaak",
	Short:   "Secure Cross Platform Application runner",
	Version: "0",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}