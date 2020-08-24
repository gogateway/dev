package main

import (
	"flag"
	"github.com/go-proxy/dev"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("port")
	if port == "" {
		p := flag.String("port", "80", "port default 80")
		port = *p
	}
	dev := dev.NewProxy()
	log.Println("start port :" + port)
	err := http.ListenAndServe(":"+port, dev)
	if err != nil {
		log.Fatal(err)
	}
}
