package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT");
	port, err := strconv.Atoi(httpPort)
    if err != nil {
   		port = 8080
    }
	http.ListenAndServe(":" + fmt.Sprint(port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, string(dump))
		log.Printf("%s", dump)
	}))
}
