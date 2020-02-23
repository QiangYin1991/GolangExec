package main

import (
	"fmt"
	"io"
	"net/http"
)



func main()  {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:30000", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(writer, "hello")
	io.WriteString(writer, " heiheihei")
}