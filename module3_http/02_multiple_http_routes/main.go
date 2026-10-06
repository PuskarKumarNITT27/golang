package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/hello",handleHello)
	http.HandleFunc("/param",handleParam)  //curl localhost:3000/param?name=puskar

	fmt.Println("server is listening at port 3000")
	http.ListenAndServe(":3000",nil)
}

func handleParam(w http.ResponseWriter ,r *http.Request){

	if r.Method != http.MethodGet {
		http.Error(w,"Only get method allowed",http.StatusMethodNotAllowed)
		return 
	}

	name := r.URL.Query().Get("name")
	
	if name == ""{
		name = "Guest"
	}

	retval := "Hello , "+ name
	_,_ = w.Write([]byte(retval))
}

func handleHello(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		http.Error(w,"Only get method allowed",http.StatusMethodNotAllowed)
		return 
	}

	_,_ = w.Write([]byte("Hello from server"))


}