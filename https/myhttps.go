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
var httpAddr = ":80"

func redirect(w http.ResponseWriter, req *http.Request) {
    // remove/add not default ports from req.Host
    target := "https://" + req.Host + req.URL.Path 
    if len(req.URL.RawQuery) > 0 {
        target += "?" + req.URL.RawQuery
    }
    log.Printf("redirect to: %s", target)
    http.Redirect(w, req, target,
            // see comments below and consider the codes 308, 302, or 301
            http.StatusPermanentRedirect)
}

func main() {
	//handle '/' route
	http.HandleFunc("/", func( res http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(res, req) 
			return
		}
		fmt.Fprint(res, "Hello World!")
	})
	go http.ListenAndServe(httpAddr, http.HandlerFunc(redirect))

	//run server on port 8443
	log.Fatal(http.ListenAndServeTLS(httpsAddr, tlsCertPath, tlsKeyPath, nil))
}