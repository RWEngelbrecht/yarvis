package helper

import (
	"encoding/json"
	"fmt"

	"github.com/RWEngelbrecht/yarvis/types"
)

func noHitResponse(resBody []byte) types.NoHitResponse {
	var apiNoHitResponse types.NoHitResponse
	err := json.Unmarshal(resBody, &apiNoHitResponse)
	if err != nil {
		OutputError(err.Error())
	}
	return apiNoHitResponse
}

func GetResponseItems(resBody []byte) types.Response {
	var apiResponse []types.ResponseItem
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

func printMeanings(meanings []types.Meaning) {
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

func OutputResponses(responseItems interface{}) {
	switch responseItems.(type) {
	case []types.ResponseItem:
		respItems := responseItems.([]types.ResponseItem)
		for _, response := range respItems {
			fmt.Println(response.Word)
			printPhonetics(response.Phonetics)
			printMeanings(response.Meanings)
			if response.Origin != "" {
				fmt.Println("Origin: " + response.Origin)
			}
		}
	case types.NoHitResponse:
		OutputError("No results found for that word...")
	default:
		OutputError("Something went wrong...")
	}
}
