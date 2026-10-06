package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main(){

	http.HandleFunc("/ok",successHandler)

	fmt.Println("Server is listening at port 3000")

	err := http.ListenAndServe(":3000",nil)

	fmt.Println(err)

}

func successHandler(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"success" :true,
		"message" : "data received",
		"datetime" : time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)

}