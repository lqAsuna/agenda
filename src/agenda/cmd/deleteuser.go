package cmd

import (
	"github.com/spf13/cobra"
)

// dmCmd represents the dm command
var deleteuserCmd = &cobra.Command{
	Use:   "dm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		password, _ := cmd.Flags().GetString("password")
		if AgendaS.DeleteUser(name, password) == true {
			fmt.Printf("Delete succeed!")
		} else {
			fmt.Printf("Delete failed!")
		}
	},
}

func init() {
	RootCmd.AddCommand(dmCmd)

	deleteuserCmd.Flags().StringP("name", "n", "Anonymous", "提供用户名字")
	deleteuserCmd.Flags().StringP("password", "n", "Anonymous", "提供用户密码")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
