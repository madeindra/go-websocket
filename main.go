package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	wsServer := newWsServer()
	go wsServer.Run()

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWS(wsServer, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
