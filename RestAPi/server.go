package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// const serverAddr string = "127.0.0.1:3000"
	http.HandleFunc("/",func (resp http.ResponseWriter, req *http.Request)  {
		fmt.Fprintln(resp, "Hello server")

		
	} )
// err:=	http.ListenAndServe(serverAddr, nil)
const PORT string =":8080"
err:=http.ListenAndServe(PORT, nil)
fmt.Println("server is listening at port no:", PORT)
if err!=nil{

	log.Fatalln("Error starting server", err)

}

}


