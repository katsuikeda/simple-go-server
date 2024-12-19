package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable must be set")
	}
	addr := ":" + port

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlePage)

	srv := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on port", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	const page = `<html>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.WriteHeader(200)
	w.Write([]byte(page))
}
