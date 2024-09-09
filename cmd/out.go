package cmd

import (
	"log"

	hrSystem "github.com/imfulee/punch/pkg/nueip"
	"github.com/spf13/cobra"
)

func CmdPunchOut() *cobra.Command {
	var nueip hrSystem.NUEIP

	punchOutCmd := &cobra.Command{
		Use:   "Out",
		Short: "Punch out",
		Long:  "Punch out NUEiP",
		Run: func(cmd *cobra.Command, args []string) {
			err := nueip.Punch(hrSystem.PunchOut)
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	punchOutCmd.Flags().StringVarP(&nueip.Username, "username", "u", "", "username of user")
	punchOutCmd.Flags().StringVarP(&nueip.Password, "password", "p", "", "password of user")
	punchOutCmd.Flags().StringVarP(&nueip.Company, "company", "c", "", "company of user")

	return punchOutCmd
}
