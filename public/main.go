package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var err error

func main() {

	fs := http.FileServer(http.Dir("."))
	if os.Getenv("GAE_ENV") == "standard" {
		http.Handle("/", firstMiddleware(fs))
	} else {
		http.Handle("/", fs)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func firstMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		hostParts := strings.Split(r.Host, ".")
		// fmt.Println(hostParts)
		// fmt.Println(r.URL.Path)
		if hostParts[0] == "www" {
			// fmt.Println("REDIREKTUJ")
			new_url := "https://" + strings.Join(hostParts[1:], ".") + r.URL.Path
			// fmt.Println(new_url)
			http.Redirect(w, r, new_url, http.StatusMovedPermanently)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
