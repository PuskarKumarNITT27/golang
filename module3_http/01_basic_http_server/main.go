package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/hello",helloHandler)	
	fmt.Println("Server is listening on port 8080")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println(err)
	}
}

func helloHandler(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w,"Only get method allowed",http.StatusMethodNotAllowed)
		return 
	}
	_, _ = w.Write([]byte("Response from server\n"))	
}