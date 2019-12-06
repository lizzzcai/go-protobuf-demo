package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/lizzzcai/go-protobuf-demo/proto/echo"
)

func makeRequest(request *echo.EchoRequest) *echo.EchoResponse {
	req, err := proto.Marshal(request)
	if err != nil {
		log.Fatalf("Unable to marshal request : %v", err)
	}

	resp, err := http.Post("http://0.0.0.0:8000/echo", "application/x-binary", bytes.NewReader(req))
	if err != nil {
		log.Fatalf("Unable to read from the server : %v", err)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Unable to read bytes from request : %v", err)
	}

	respObj := &echo.EchoResponse{}
	err = proto.Unmarshal(respBytes, respObj)
	if err != nil {
		log.Fatalf("Unable to un-marshal bytes to object : %v", err)
	}

	return respObj
}

func main() {
	request := &echo.EchoRequest{Name: "lizzzcai"}
	resp := makeRequest(request)
	log.Printf("Response from API is : %v\n", resp.GetMessage())
}