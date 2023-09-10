package cmd

import (
	"github.com/imfulee/punch/hr_system"
	"github.com/spf13/cobra"
)

func CmdPunchOut() *cobra.Command {
	var nueip hr_system.NUEIP

	punchOutCmd := &cobra.Command{
		Use:              "in",
		Short:            "Punch in",
		Long:             "Punch in NUEiP",
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			nueip.Punch(hr_system.PunchOut)
		},
	}

	punchOutCmd.Flags().StringVarP(&nueip.Username, "username", "u", "", "username of user")
	punchOutCmd.Flags().StringVarP(&nueip.Password, "password", "p", "", "password of user")
	punchOutCmd.Flags().StringVarP(&nueip.Company, "company", "c", "", "company of user")

	return punchOutCmd
}
