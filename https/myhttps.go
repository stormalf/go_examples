package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var tlsCertPath = os.Getenv("TLS_CERT_PATH")
var tlsKeyPath = os.Getenv("TLS_KEY_PATH")

func httpsRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
}

func main() {
	//handle '/' route
	http.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})
	go http.ListenAndServe(":81", http.HandlerFunc(httpsRedirect))

	//run server on port 443
	log.Fatal(http.ListenAndServeTLS(":8443", tlsCertPath, tlsKeyPath, nil))
}