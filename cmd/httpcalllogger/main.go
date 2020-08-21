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
	err := http.ListenAndServe(":10080", nil)
	if err != nil {
		panic(err)
	}
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

