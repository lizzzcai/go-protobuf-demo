package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/lizzzcai/go-protobuf-demo/proto/echo"
)

func Echo(writer http.ResponseWriter, req *http.Request) {
	contentLength := req.ContentLength
	log.Printf("Content Length Received: %v\n", contentLength)
	request := &echo.EchoRequest{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Unable to read message from request : %v", err)
	}

	err = proto.Unmarshal(data, request)
	if err != nil {
		log.Fatalf("Unable to un-marshal the message : %v", err)
	}

	name := request.GetName()
	responseContent := &echo.EchoResponse{Message: fmt.Sprintf("Hello %s", name)}
	response, err := proto.Marshal(responseContent)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	_, err = writer.Write(response)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Println("Starting the API server...")
	r := mux.NewRouter()
	r.HandleFunc("/echo", Echo).Methods("POST")

	server := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
