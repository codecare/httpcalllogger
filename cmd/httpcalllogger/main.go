package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("started")

	http.HandleFunc("/", logcall)
	http.ListenAndServe(":10080", nil)
}

func logcall(w http.ResponseWriter, req *http.Request) {

	f, err := os.Create("header.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	fmt.Fprint(f, "requestURI: " + req.RequestURI + "\n")
	fmt.Fprint(f, "requestMethod: " + req.Method + "\n")
	header := req.Header
	header.Write(f)

	all, _ := ioutil.ReadAll(req.Body)

	err = ioutil.WriteFile("data.bin", all, 0644)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("logged request")
}

