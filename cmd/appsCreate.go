package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// appsCreateCmd represents the appsCreate command
var appsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new App",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("appsCreate called")
	},
}

func init() {
	appsCmd.AddCommand(appsCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appsCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appsCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
