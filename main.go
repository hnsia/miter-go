package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Input struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Output struct {
	A int `json:"a"`
	B int `json:"b"`
	Sum int `json:"sum"`
}

func sum(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	var i Input
	json.Unmarshal(b, &i)
	log.Println(string(b))
	var o Output
	o.A = i.A
	o.B = i.B
	o.Sum = i.A + i.B
	b, _ = json.Marshal(o)
	fmt.Fprint(w, string(b))
}

func main() {
	srv := http.NewServeMux()
	srv.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `Hello World!`)
	})
	srv.HandleFunc("POST /sum", sum)
	http.ListenAndServe(":8080", srv)
	
}