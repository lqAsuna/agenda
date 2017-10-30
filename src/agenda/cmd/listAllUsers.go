package cmd

import (
	"agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

// dmCmd represents the dm command
var listAllUsersCmd = &cobra.Command{
	Use:   "listAllUsers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		qr := AgendaS.QueryAllUsers()
		for e := qr.Front(); e != nil; e = e.Next() {
			tur := e.Value.(entity.User)
			fmt.Printf("Name: %v\n", tur.Name)
			fmt.Printf("Email: %v\n", tur.Email)
			fmt.Printf("Phone: %v\n\n", tur.Phone)
		}
	},
}

func init() {
	RootCmd.AddCommand(listAllUsersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
