package main

import "fmt"

func ListenAndServe(addr string, handler Handler) error

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

func main() {
	fmt.Println("vim-go")
}
