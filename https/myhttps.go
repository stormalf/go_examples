package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var tlsCertPath = os.Getenv("TLS_CERT_PATH")
var tlsKeyPath = os.Getenv("TLS_KEY_PATH")
var httpsAddr = ":8443"


func main() {
	//handle '/' route
	http.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(res, req) 
			return
		}
		fmt.Fprint(res, "Hello World!")
	})

	//run server on port 443
	log.Fatal(http.ListenAndServeTLS(httpsAddr, tlsCertPath, tlsKeyPath, nil))
}