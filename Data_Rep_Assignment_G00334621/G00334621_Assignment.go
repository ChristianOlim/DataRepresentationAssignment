package main

import(
	"fmt"
	"net/http"
)

func main(){
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
	//fmt.Println("Just trying a template.")
}