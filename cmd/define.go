/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	h "github.com/RWEngelbrecht/yarvis/helper"
	t "github.com/RWEngelbrecht/yarvis/types"
	"github.com/spf13/cobra"
)

// defineCmd represents the define command
var defineCmd = &cobra.Command{
	Use:   "define",
	Short: "Define a given word.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			h.OutputError("You need to give me a single word, Human...")
		}

		lookup_word := args[0]
		// do some word validation

		body := h.LookupWord(lookup_word)

		var responseItems t.Response = h.GetDefinitionResponseItems(body)
		h.OutputResponses(responseItems)
	},
}

func init() {
	rootCmd.AddCommand(defineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// defineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// defineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
