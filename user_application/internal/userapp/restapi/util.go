package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, status int, response any) error {
	jsonString, err := json.Marshal(response)
	if err != nil {
		return fmt.Errorf("json marshal response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(jsonString)
	if err != nil {
		return fmt.Errorf("write response body: %w", err)
	}

	return nil
}
