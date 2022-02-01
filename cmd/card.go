/*
Copyright © 2022 Sam Kim <netpple@gmail.com>
This file is part of Sam's CLI myapp.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "This command creates awesome greetings",
	Long: ` For example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr
	greetctl create card bob --name="Bob Marley" --occasion=birthday`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		fmt.Println("value of the flag name: "+ name)
		fmt.Println("value of the flag language: " + language)
	},
}

func init() {
	createCmd.AddCommand(cardCmd)
	cardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	cardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occasion")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
