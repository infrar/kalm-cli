package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routesCreateCmd represents the routesCreate command
var routesCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new certificate via Let's Encrypt",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("routesCreate called")
	},
}

func init() {
	routesCmd.AddCommand(routesCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routesCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routesCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
