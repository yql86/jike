package main

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net/http"
)

func main() {
	glog.Infoln("enter main...")
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("entering root handler")
	user := request.URL.Query().Get("user")
	if user != "" {
		io.WriteString(writer, fmt.Sprintf("hello[%s]\n", user))
	} else {
		io.WriteString(writer, "hello [stranger]\n")
	}
}
