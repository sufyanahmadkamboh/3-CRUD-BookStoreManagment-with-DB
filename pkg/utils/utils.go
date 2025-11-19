package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParsBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}
