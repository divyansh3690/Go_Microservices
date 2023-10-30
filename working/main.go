package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"working/handlers"
)

func main() {
	//both the handlers 
	customHandler := handlers.NewHomePage()
	goodbyeHandler := handlers.Caller_func()

//  custom servMux
	servmux := http.NewServeMux()
	servmux.Handle(("/goodbye"), goodbyeHandler)
// handle func
	servmux.Handle("/", customHandler)
	
	
	log.Println("Welcome to server")
	srv := &http.Server{
		Addr:        ":9000",
		Handler:     servmux,
		IdleTimeout: 120 * time.Second,
		//Idle Timeout-->max amt of time that an idle connection can remain open
		ReadTimeout: 30 * time.Second,
		// max time server will wait for any data to be sent by the client.
		WriteTimeout: 1 * time.Second,
		//
	}

	// concurrently runs server and checks interruptions

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	// gracefully shutdown server
	 
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Print("recieved terminate, sever gonna shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)

}

// -->This code contains listen and server ie., a http function we can go more bare metal and
// customize it to our needs

// func main() {

// 	customHandler:=handlers.NewHomePage()

// 	servmux:=http.NewServeMux()
// 	goodbyeHandler:=handlers.Caller_func()
// 	servmux.Handle("/",customHandler)
// 	servmux.Handle(("/goodbye"),goodbyeHandler)
// 	log.Println("Welcome to server")
//     http.ListenAndServe(":9000", servmux)
// }

// --> this is code that uses default at every point has a
// default server multiplexer handler etc
// //package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	fmt.Println("Hello")
// 	http.HandleFunc("/", func(response http.ResponseWriter, req *http.Request) {
// 		log.Println("Welcome to homepage")
// 		data, err := io.ReadAll(req.Body)
// 		if err != nil {
// 			response.WriteHeader(http.StatusBadGateway)
// 			response.Write([]byte("Opps wrong "))
// 			return
// 		}
// 		fmt.Println("Hello dfusnoHAIXOHNXOANFUD  9WA",data)

// 	})

// 	http.ListenAndServe(":9000", nil)
// }

// func Error_check(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }



// http.listenAndServe(Port , srevmux)


// 