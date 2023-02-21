package helper

import (
	"encoding/json"
)

func GetLanguage(language, data string) (string, error) {
	var model map[string]string
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		return "", err
	}

	return model[language], err
}
