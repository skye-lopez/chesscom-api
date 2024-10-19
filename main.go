package chesscom

import (
	"encoding/json"
	"net/http"
)

func request(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}
