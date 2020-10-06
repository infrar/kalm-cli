package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// routesUploadCmd represents the routesUpload command
var routesUploadCmd = &cobra.Command{
	Use:   "upload -f [path/to/file]",
	Short: "Upload an existing certificate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Cert Uploaded")
	},
}

func init() {
	certsCmd.AddCommand(routesUploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routesUploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routesUploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
