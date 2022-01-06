/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	h "github.com/RWEngelbrecht/yarvis/helper"
	"github.com/spf13/cobra"
)

// countdownCmd represents the countdown command
var countdownCmd = &cobra.Command{
	Use:     "countdown",
	Aliases: []string{"cd", "c"},
	Short:   "yarvis countdown year/month/day hh:mm:ss",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 || len(args) == 0 {
			h.OutputError("You need to give me a date to count down to, Human...")
		}
		// calculate until second, if given
		// calculate until minute, if given
		// calculate until midnight if only day is given
		until := h.ParseDateTime(args)
		fmt.Println(until)
	},
}

func init() {
	rootCmd.AddCommand(countdownCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countdownCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countdownCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
