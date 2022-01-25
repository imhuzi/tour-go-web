package main

import (
	"fmt"
	"log"
	"net/http"
)

// main 函数
func main() {
	fmt.Print("Start Web Server" );
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi there , I love %s!", r.URL.Path[1:])
}