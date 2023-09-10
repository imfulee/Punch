package cmd

import (
	"github.com/imfulee/punch/hr_system"
	"github.com/spf13/cobra"
)

func CmdPunchIn() *cobra.Command {
	var nueip hr_system.NUEIP

	punchInCmd := &cobra.Command{
		Use:   "In",
		Short: "Punch in",
		Long:  "Punch in NUEiP",
		Run: func(cmd *cobra.Command, args []string) {
			nueip.Punch(hr_system.PunchIn)
		},
	}

	punchInCmd.Flags().StringVarP(&nueip.Username, "username", "u", "", "username of user")
	punchInCmd.Flags().StringVarP(&nueip.Password, "password", "p", "", "password of user")
	punchInCmd.Flags().StringVarP(&nueip.Company, "company", "c", "", "company of user")

	return punchInCmd
}
