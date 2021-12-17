package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	c "github.com/RWEngelbrecht/yarvis/colours"
	"github.com/RWEngelbrecht/yarvis/types"
)

func OutputError(errText string) {
	l := log.New(os.Stderr, "", 0)
	l.Println(string(c.Red) + errText)
	// panic(string(c.Red), errors.New(errText).Error())
}

func noHitResponse(resBody []byte) types.NoHitResponse {
	var apiNoHitResponse types.NoHitResponse
	err := json.Unmarshal(resBody, &apiNoHitResponse)
	if err != nil {
		OutputError(err.Error())
	}
	return apiNoHitResponse
}

func OutputResponses(responseItems interface{}) {
	switch responseItems.(type) {
	case []types.DefinitionResponseItem:
		respItems := responseItems.([]types.DefinitionResponseItem)
		for _, response := range respItems {
			fmt.Println(response.Word)
			printPhonetics(response.Phonetics)
			printDefinitionMeanings(response.Meanings)
			if response.Origin != "" {
				fmt.Println("Origin: " + response.Origin)
			}
		}
	case []types.SynonymResponseItem:
		respItems := responseItems.([]types.SynonymResponseItem)
		for _, response := range respItems {
			fmt.Println(response.Word)
			printSynonymMeanings(response.Meanings)
		}
	case types.NoHitResponse:
		OutputError("No results found for that word...")
	default:
		OutputError("Something went wrong...")
	}
}
