package lib

import (
	"net/http"
	"encoding/json"
	"errors"
	"io"
)

func JsonDecodeRequestBody(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)
	ErrorCheck(err)
}

func JsonDecodeFromStream(body io.Reader) (interface{}, error) {
	var v interface{}

	if err := json.NewDecoder(body).Decode(&v); err != nil {
		return nil, errors.New("unknown json unmarshal error")
	}

	return v, nil
}
