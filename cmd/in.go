package cmd

import (
	"fmt"

	hrSystem "github.com/imfulee/punch/pkg/nueip"
	"github.com/spf13/cobra"
)

func CmdPunchIn() *cobra.Command {
	var nueip hrSystem.NUEIP

	punchInCmd := &cobra.Command{
		Use:   "In",
		Short: "Punch in",
		Long:  "Punch in NUEiP",
		Run: func(cmd *cobra.Command, args []string) {
			err := nueip.Punch(hrSystem.PunchIn)
			if err != nil {
				fmt.Println(err)
			}
		},
	}

	punchInCmd.Flags().StringVarP(&nueip.Username, "username", "u", "", "username of user")
	punchInCmd.Flags().StringVarP(&nueip.Password, "password", "p", "", "password of user")
	punchInCmd.Flags().StringVarP(&nueip.Company, "company", "c", "", "company of user")

	return punchInCmd
}
