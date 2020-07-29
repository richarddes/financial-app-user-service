// Package internal provides common globally used functions.
package internal

import (
	"encoding/json"
	"io"
	"user-service/config"
)

// ParseJSONBody decodes a jsonBody and saves it in the value pointed by v.
func ParseJSONBody(jsonBody io.ReadCloser, body interface{}) error {
	dc := json.NewDecoder(jsonBody)
	err := dc.Decode(body)
	if err != nil {
		return err
	}

	return nil
}

// IsSupportedLang checks if the specified language is in the config's SupportedLangs array
func IsSupportedLang(lang string) bool {
	for _, i := range config.SupportedLangs {
		if lang == i {
			return true
		}
	}

	return false
}
