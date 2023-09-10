package cmd

import "github.com/spf13/cobra"

func CmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "punch",
		Short: "Punch punches card",
		Long:  "Punch punches card for NUEiP",
	}

	rootCmd.AddCommand(CmdPunchIn())
	rootCmd.AddCommand(CmdPunchOut())

	return rootCmd
}
