// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// chptrCmd represents the chptr command
var chptrCmd = &cobra.Command{
	Use:   "chptr",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("participators")
		del, _ := cmd.Flags().GetBool("delete")
		add, _ := cmd.Flags().GetBool("add")

		if del {
			//
			AgendaS.DeelteMeetingParticipators(AgendaS.GetAgendaServiceStorage().Current.Name, title, participators)
		} else if add {
			AgendaS.AddMeetingParticipators(AgendaS.GetAgendaServiceStorage().Current.Name, title, participators)
		}
	},
}

func init() {
	RootCmd.AddCommand(chptrCmd)

	chptrCmd.Flags().StringP("title", "t", "Anonymous", "提供会议主题名字")
	chptrCmd.Flags().StringArrayP("participators", "p", nil, "提供会议的参与者")
	chptrCmd.Flags().BoolP("delete", "d", false, "删除会议的某些参与者")
	chptrCmd.Flags().BoolP("add", "a", false, "删除会议的某些参与者")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chptrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chptrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
