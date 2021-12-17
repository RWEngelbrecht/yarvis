/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"
	"net/http"

	h "github.com/RWEngelbrecht/yarvis/helper"
	t "github.com/RWEngelbrecht/yarvis/types"
	"github.com/spf13/cobra"
)

// synonymCmd represents the synonym command
var synonymCmd = &cobra.Command{
	Use:   "synonym",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			h.OutputError("You need to give me a single word, Human...")
		}

		lookup_word := args[0]
		// do some word validation

		endpoint := "https://api.dictionaryapi.dev/api/v2/entries/en_GB/" + lookup_word

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

		var responseItems t.Response = h.GetSynonymResponseItems(body)
		h.OutputResponses(responseItems)
	},
}

func init() {
	rootCmd.AddCommand(synonymCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// synonymCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// synonymCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
