package lib

import (
	"net/http"
	"encoding/json"
	"log"
)

func JsonDecode(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)
	if err != nil {
		log.Fatal(err)
	}
}