package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
	Length int `json:"length"`
}
func main(){

	url := `https://catfact.ninja/fact`

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	bodyBytes,err := io.ReadAll(res.Body)
	
	if err != nil {
		fmt.Println("reading body failed ",err)
		return
	}

	var data CatFactResponse

	if err:= json.Unmarshal(bodyBytes,&data) ; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.Fact, data.Length)
}