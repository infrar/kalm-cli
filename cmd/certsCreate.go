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
	"time"

	"github.com/janeczku/go-spinner"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// certsCreateCmd represents the certsCreate command
var certsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Certificate",
	Run: func(cmd *cobra.Command, args []string) {
		createCert()
	},
}

func createCert() {
	var err error

	// 1. name
	prompt := promptui.Prompt{
		Label:   "Domain Name",
		Default: "myapp.kapp.live",
	}

	certName, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// 2. http/https

	prompt2 := promptui.Select{
		Label: "Protocol",
		Items: []string{"Http", "Https"},
	}

	_, protocolType, err := prompt2.Run()

	// 3. provision
	s := spinner.StartNew("Provisioning Certificate...")
	time.Sleep(2 * time.Second) // something more productive here
	s.Stop()

	fmt.Printf("%s Certificate for `%s` provisioned\n", protocolType, certName)
}

func init() {
	certsCmd.AddCommand(certsCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// certsCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// certsCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
