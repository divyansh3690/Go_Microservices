package handlers

import (
	"fmt"
	"io"
	"net/http"
)

type custom struct {
}

func Caller_func() (*custom){
	return &custom{}
}




func (*custom) ServeHTTP(rw http.ResponseWriter,req*http.Request){
	data,err:=io.ReadAll(req.Body)
	if err!=nil{
		rw.WriteHeader(http.StatusBadGateway)
		rw.Write([]byte("OPPPSSS"))
		return
	}

	fmt.Fprintf(rw,"Hello from Server:%s",data)

}
