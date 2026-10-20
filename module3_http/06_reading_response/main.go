package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todos struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main(){
	url := `https://jsonplaceholder.typicode.com/todos`

	res,err := http.Get(url)

	if err != nil {
		fmt.Println(`Error in fetching data`)
	}
	
	defer res.Body.Close()

	if res.StatusCode  != http.StatusOK {
		fmt.Println(res.Status)
		return 
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return 
	}

	// bodyText := string(bodyBytes)

	// only reading the json docs
	// max := 250

	// if(len(bodyText) < max){
	// 	max = len(bodyText)
	// }

	// fmt.Println(bodyText[:max])

	// unmarshalling json 

	var todos[] Todos

	if err := json.Unmarshal(bodyBytes,&todos) ; err != nil {
		fmt.Println(err)
		return 
	}

	for _,todo := range todos {
		fmt.Println("Todos No: ",todo.Id)
		fmt.Println(todo.Title)
		fmt.Println(todo.Completed)

	}
}