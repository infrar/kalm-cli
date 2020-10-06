package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// certsDeleteCmd represents the certsDelete command
var certsDeleteCmd = &cobra.Command{
	Use:   "delete [certificate name]",
	Short: "Delete a certificate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Certificate Deleted")
	},
}

func init() {
	certsCmd.AddCommand(certsDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// certsDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// certsDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
