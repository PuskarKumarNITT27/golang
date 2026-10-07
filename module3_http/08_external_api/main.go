package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func writeJSON(w http.ResponseWriter, status int , data any){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFactResponse,error){
	url := `https://catfact.ninja/fact`
	res,err := http.Get(url)

	if err != nil {
		return CatFactResponse{},err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CatFactResponse{} ,fmt.Errorf("external api failed :%s",res.Status)
	}

	bodyBytes ,err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return CatFactResponse{} ,err 
	}

	var data CatFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil{
		return CatFactResponse{},err
	}

	return data, nil
}
func externalHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		writeJSON(w,http.StatusMethodNotAllowed,map[string]any{
			"success":false,
			"error":"Only get method allowed",
		})
		return 
	}

	data , err := fetchCatFact()

	if err != nil {
		writeJSON(w,http.StatusBadGateway,map[string]any{
			"success":false,
			"error":"failed to fetch data",
		})
		return 
	}

	writeJSON(w,http.StatusOK,map[string]any{
		"success":true,
		"timestamp":time.Now().UTC(),
		"external": map[string]any{
			"source": "catfact.ninja",
			"fact":data.Fact,
			"length":data.Length,
		},
	})


}

func main(){
	http.HandleFunc("/external",externalHandler)
	err := http.ListenAndServe(":3000",nil)
	
	fmt.Println(err)
}