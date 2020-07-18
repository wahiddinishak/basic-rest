package main

import (
	"fmt"
	"net/http"
	"os"
)

func setResp(res http.ResponseWriter, msg []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(msg)
}

func main() {

	// main route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		msg := []byte(`"{message}":"Server is up!"`)
		setResp(res, msg, http.StatusOK)
	})

	err := http.ListenAndServe(":65", nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
