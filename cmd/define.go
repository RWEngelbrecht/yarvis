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
	"fmt"
	"io/ioutil"
	"net/http"

	h "github.com/RWEngelbrecht/yarvis/helper"
	"github.com/spf13/cobra"
)

func print_definitions() {

}

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

		endpoint := "https://api.dictionaryapi.dev/api/v2/entries/en_GB/" + lookup_word
		fmt.Println(endpoint)
		res, err := http.Get(endpoint)
		if err != nil {
			h.OutputError("Something went wrong with the request...")
			// try another api if this fails
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			h.OutputError("Something went wrong when reading the response body...")
		}

		fmt.Println(body[0])
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