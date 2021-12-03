package types

type Response interface{}

type PhoneticItem struct {
	Text  string `json:"text"`
	Audio string `json:"audio"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

// add structs for definition api response
type ResponseItem struct {
	Word      string         `json:"word"`
	Phonetic  string         `json:"phonetic"`
	Phonetics []PhoneticItem `json:"phonetics"`
	Origin    string         `json:"origin"`
	Meanings  []Meaning      `json:"meanings"`
}

type NoHitResponse struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}
