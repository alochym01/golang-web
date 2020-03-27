package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL)
	fmt.Println(req.URL.Scheme)
	fmt.Fprintf(res, "Hello\n")
}

func main() {

	// using sayHello to handle request
	http.HandleFunc("/", sayHello)

	// using embeded function to handle request
	http.HandleFunc("/alochym", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
		fmt.Println(req.URL.Scheme)
		fmt.Fprintf(res, "Hello Alochym\n")
	})

	// starting http server w port 500
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal("listen ans serve: ", err)
	}
	fmt.Println("server running")
}
