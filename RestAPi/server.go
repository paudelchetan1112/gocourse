// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	// const serverAddr string = "127.0.0.1:3000"
// 	http.HandleFunc("/",func (resp http.ResponseWriter, req *http.Request)  {
// 		fmt.Fprintln(resp, "Hello server")

// 	} )
// // err:=	http.ListenAndServe(serverAddr, nil)
// const PORT string =":8080"
// err:=http.ListenAndServe(PORT, nil)
// fmt.Println("server is listening at port no:", PORT)
// if err!=nil{

// 	log.Fatalln("Error starting server", err)

// }

// }

package main

import (
	"fmt"
	"io"
	"net/http"
)


func main(){

	//create a new http client

	client:=&http.Client{

	}
	resp, err:=client.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err!=nil{
		fmt.Println("Error making get request", err)
	}
		defer resp.Body.Close()
		//Read and print the response body
		body, err:=io.ReadAll(resp.Body)

		if err!=nil{
			fmt.Println("Error readin response body:", err)
			return 
		}
		fmt.Println(body)
		fmt.Println(string(body))

	



}