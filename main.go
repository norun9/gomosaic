package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	TITLEDB = tilesDB()
	fmt.Println("Mosaic server started")
	server.ListenAndServe()
}
