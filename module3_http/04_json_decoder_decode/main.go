package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type TestRequest struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/test", testHandler)

	fmt.Println("server is listening on port 3000")
	err := http.ListenAndServe(":3000",nil)

	fmt.Println(err)
}

func writeJSON(w http.ResponseWriter, status int , data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func testHandler(w http.ResponseWriter , r *http.Request){

	if r.Method != http.MethodPost {
		writeJSON(w,http.StatusMethodNotAllowed,map[string]any{
			"success":false,
			"error" : "method not allowed",
		})
		return 
	}

	defer r.Body.Close()

	var req TestRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req) ; err != nil {
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"success":false,
			"error":"Incorrect JOSN format",
		})
		return 
	}

	req.Name = strings.TrimSpace(req.Name)

	if req.Name == ""{
		writeJSON(w,http.StatusBadRequest,map[string]any{
			"success":false,
			"error":"Name should not be empty",
		})
		return 
	}

	writeJSON(w,http.StatusOK,map[string]any{
		"success":true,
		"data":req.Name,
		"timestamp":time.Now().UTC(),
	})

}