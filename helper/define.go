package helper

import (
	"encoding/json"
	"fmt"

	"github.com/RWEngelbrecht/yarvis/types"
)

func GetDefinitionResponseItems(resBody []byte) types.Response {
	var apiResponse []types.DefinitionResponseItem
	err := json.Unmarshal(resBody, &apiResponse)
	if err != nil {
		return noHitResponse(resBody)
	}
	return apiResponse
}

func printDefinitions(definitions []types.Definition) {
	for _, definition := range definitions {
		fmt.Println(" - " + definition.Definition)
	}
}

func printDefinitionMeanings(meanings []types.Meaning) {
	for _, meaning := range meanings {
		fmt.Println("PoS: " + meaning.PartOfSpeech)
		fmt.Println("Definitions:")
		printDefinitions(meaning.Definitions)
	}
}

func printPhonetics(phonetics []types.PhoneticItem) {
	for _, phonetic := range phonetics {
		if phonetic.Text != "" {
			fmt.Println("( " + phonetic.Text + " )")
		}
		fmt.Println("Audio: " + phonetic.Audio)
	}
}
