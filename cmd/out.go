package cmd

import (
	"fmt"

	"github.com/imfulee/punch/hr_system"
	"github.com/spf13/cobra"
)

func CmdPunchOut() *cobra.Command {
	var nueip hr_system.NUEIP

	punchOutCmd := &cobra.Command{
		Use:   "Out",
		Short: "Punch out",
		Long:  "Punch out NUEiP",
		Run: func(cmd *cobra.Command, args []string) {
			err := nueip.Punch(hr_system.PunchOut)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	punchOutCmd.Flags().StringVarP(&nueip.Username, "username", "u", "", "username of user")
	punchOutCmd.Flags().StringVarP(&nueip.Password, "password", "p", "", "password of user")
	punchOutCmd.Flags().StringVarP(&nueip.Company, "company", "c", "", "company of user")

	return punchOutCmd
}
