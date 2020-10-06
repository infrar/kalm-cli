/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appsListCmd represents the appsList command
var appsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Apps",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Kalm Apps:\n\n")

		for i, app := range fakeApps {
			fmt.Printf("%d) %s\n", i+1, app)
		}
	},
}

func init() {
	appsCmd.AddCommand(appsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
