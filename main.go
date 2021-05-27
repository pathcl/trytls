package main

import (
	"flag"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {

	log.Println("App Started")
	domain := flag.String("domain", "test.domain.tld/", "domain to test")
	cert := flag.String("cert", "cert.pem", "public key to use")
	key := flag.String("key", "cert.key", "private key to use")

	flag.Parse()

	http.HandleFunc(*domain, HelloServer)
	err := http.ListenAndServeTLS(":8443", *cert, *key, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
