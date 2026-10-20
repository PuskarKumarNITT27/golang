package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFactResponse, error) {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		return CatFactResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CatFactResponse{}, fmt.Errorf("external api failed: %s", res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return CatFactResponse{}, err
	}

	var data CatFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return CatFactResponse{}, err
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Only GET method is allowed",
		})
		return
	}

	data, err := fetchCatFact()
	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": "Failed to fetch data",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":        true,
		"timeStamp": time.Now().UTC(),
		"external": map[string]any{
			"source": "Catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})
}

func main() {

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)
}
