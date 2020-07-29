package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
	"user-service/config"
)

func toJSON(t *testing.T, i interface{}) io.ReadCloser {
	t.Helper()

	b, err := json.Marshal(i)
	if err != nil {
		t.Fatal(err)
	}

	rc := ioutil.NopCloser(bytes.NewReader(b))

	return rc
}

func TestParseJSONBody(t *testing.T) {
	var body config.StockBody

	content := config.StockBody{Symbol: "AAPL", Amount: 105, Price: 250.03}
	err := ParseJSONBody(toJSON(t, content), &body)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(content, body) {
		t.Errorf("Expected parsed body to equal %v but got %v instead", content, body)
	}
}
