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
