package lib

import (
	"net/http"
	"encoding/json"
)

func JsonDecodeRequestBody(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)
	ErrorCheck(err)
}