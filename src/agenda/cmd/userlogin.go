package cmd

import (
	"agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

var userloginCmd = &cobra.Command {
	Use:   "userlogin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		ur := entity.GetUser(name, password, "", "")
		if AgendaS.UserLogIn(ur) == true {
			fmt.Printf("Log In Succeed!\n")
			AgendaS.GetAgendaServiceStorage().WriteCurUsr("entity/curUser.txt")
		} else {
			fmt.Printf("Log In Failed!\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(userloginCmd)
	userloginCmd.Flags().StringP("name", "n", "Anonymous", "提供用户名字")
	userloginCmd.Flags().StringP("password", "p", "Anonymous", "提供用户密码")
		// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
