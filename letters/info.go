package letters

import (
	"encoding/json"
	"net/http"
)

type WordInfo struct {
	Word     string
	Meanings []struct {
		Definitions []struct {
			Definition string `json:"definition"`
		}
	}
}

func GetWordInfo(word string) (WordInfo, bool) {
	res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + word)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var wordInfo []WordInfo
	err = json.NewDecoder(res.Body).Decode(&wordInfo)
	if err != nil || len(wordInfo) == 0 || wordInfo[0].Word == "" {
		return WordInfo{}, false
	}

	return wordInfo[0], true
}

func (w WordInfo) String() string {
	var meanings []string
	for _, meaning := range w.Meanings {
		for _, definition := range meaning.Definitions {
			meanings = append(meanings, definition.Definition)
		}
	}
	return w.Word + ": " + meanings[0]
}
