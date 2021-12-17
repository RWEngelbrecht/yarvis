package helper

import (
	"encoding/json"
	"fmt"

	"github.com/RWEngelbrecht/yarvis/types"
)

func GetSynonymResponseItems(resBody []byte) types.Response {
	var apiResponse []types.SynonymResponseItem
	err := json.Unmarshal(resBody, &apiResponse)
	if err != nil {
		return noHitResponse(resBody)
	}
	return apiResponse
}

func printSynonyms(definitions []types.Definition) {
	for _, definition := range definitions {
		for i, synonym := range definition.Synonyms {
			switch {
			case i == len(definition.Synonyms)-1:
				fmt.Println(synonym)
			case i%5 == 0 && i != 0:
				fmt.Println()
				fallthrough
			default:
				fmt.Print(synonym + ", ")
			}
		}
		fmt.Println()
	}
}

func printSynonymMeanings(meanings []types.Meaning) {
	for _, meaning := range meanings {
		fmt.Println("PoS: " + meaning.PartOfSpeech)
		fmt.Println("Synonyms:")
		printSynonyms(meaning.Definitions)
	}
}
