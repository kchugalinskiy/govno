package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
{
	"name": "vasya",
	"age": 12,
	"address": {
		"city": "Moscow"
	},
	"labels":["good guy", "best friend"]
}
*/

type Address struct {
	City string `json:"city"`
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
	Labels  []string `json:"labels"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Person{
			Name: "vasya",
			Age:  12,
			Address: &Address{
				City: "Moscow",
			},
			Labels: []string{"good guy", "best friend"},
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(p)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body><h1>Hello</h1></body></html>"))
	})
	http.ListenAndServe(":8080", nil)
}
