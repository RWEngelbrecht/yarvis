package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	f "github.com/RWEngelbrecht/yarvis/flavour"
	"github.com/RWEngelbrecht/yarvis/types"
)

func OutputError(errText string) {
	l := log.New(os.Stderr, "", 0)
	l.Println(string(f.Red) + errText)
	// panic(string(f.Red), errors.New(errText).Error())
}

func LookupWord(word string) []byte {
	endpoint := "https://api.dictionaryapi.dev/api/v2/entries/en_GB/" + word
	req, _ := http.NewRequest("GET", endpoint, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		OutputError(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		OutputError(err.Error())
	}
	return body
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
