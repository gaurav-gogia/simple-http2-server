package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./*.html"))
}

func main() {
	http.Handle("/assets/",
		http.StripPrefix("/assets",
			http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", index)

	go http.ListenAndServe(":5000", nil)
	err := http.ListenAndServeTLS(":8888", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {	
	if pusher, ok := w.(http.Pusher); ok {
		options := &http.PushOptions{
			Header: http.Header{
				"Accept-Encoding": r.Header["Accept-Encoding"],
			},
		}
		if err := pusher.Push("/assets/js/login.js", options); err != nil {
			fmt.Println(err)
		}
		if err := pusher.Push("/assets/css/normalizeLogin.css", options); err != nil {
			fmt.Println(err)
		}
		if err := pusher.Push("/assets/css/styleLogin.css", options); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("COULD NOT PUSH")
	}

	tpl.ExecuteTemplate(w, "cook.html", nil)
	log.Println(r.URL.Path)
}

/*
   Process for handling TLS Secure Connections for your WebServices:
   1. go run C:\Go\src\crypto\tls\generate_cert.go --host=<domain_name>
      Do that to generate certificate & key or use StartSSL or something.
   2. Hand-over your certificate to another company which will sign them.

   Otherwise you will have to go to some Certificate Authority and get
   your Certificate signed.

   Cloudflare provides free SSL / TLS, self-signed certificates work there

   One of the free companies is https://letsencrypt.org/
*/
