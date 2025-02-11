package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		log.Println("Request body is empty")
		return
	}

	if err := json.Unmarshal(body, x); err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}
}
