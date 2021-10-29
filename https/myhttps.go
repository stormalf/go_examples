package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tlsCertPath = os.Getenv("TLS_CERT_PATH")
var tlsKeyPath = os.Getenv("TLS_KEY_PATH")
var httpsAddr = ":8443"
var httpAddr = ":8082"
var tmpl *template.Template
var login = "login.html"

func init() {
    tmpl = template.Must(template.ParseFiles("template/"+login))
}

//redirect http to https
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
		tmpl.ExecuteTemplate(res, login, nil)
	})
	go func() {
		if err := http.ListenAndServe(httpAddr, http.HandlerFunc(redirect)); err != nil {
			log.Fatalf("http error: %v", err)
		}
	}()
	//go http.ListenAndServe(httpAddr, http.HandlerFunc(redirect))

	//run server on port 443
	log.Fatal(http.ListenAndServeTLS(httpsAddr, tlsCertPath, tlsKeyPath, nil))
}