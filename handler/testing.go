package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func toJSON(body interface{}) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// NewRecorder returns a new ResponseRecorder. It sets a a uid and lang header.
func NewRecorder(t *testing.T, method, url string, body interface{}, handler http.HandlerFunc) *httptest.ResponseRecorder {
	t.Helper()

	jsonBody, err := toJSON(body)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Add("uid", strconv.FormatUint(1, 10))
	req.Header.Add("lang", "en")

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)

	h.ServeHTTP(rr, req)

	return rr
}
