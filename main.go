package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const vidsDir = "videos"
	const port = 8069

	http.Handle("/", addHeaders(http.FileServer(http.Dir(vidsDir))))
	log.Printf("Serving %s on HTTP port: %v\n", vidsDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
