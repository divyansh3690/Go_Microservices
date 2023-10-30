package handlers

import (
	"fmt"
	"io"
	"net/http"
)

type homepage struct{
	
}

//return a pointer of homepage structure
func NewHomePage() (*homepage) {
	return &homepage{}
}


//ServerHTTP func is associated with pointer to an object of homepage struct and ServeHTTP is also a
// function defined in Handler interface that is passed in Handle function .
// Thus we need to pass any type that has a function defined as ServeHTTP 
func (* homepage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		rw.Write([]byte("Wrong OUTPUT"))
		return
	}
	// fmt.Print(data)
	fmt.Fprintf(rw, "Hello, World! %s", data)
}
