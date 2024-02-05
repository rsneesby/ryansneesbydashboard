package main

import (
	"fmt"
	"net/http"
)

// Middleware to enable CORS
func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Explicitly specify allowed methods
		// Explicitly specify allowed headers, including 'hx-request'
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, hx-request, hx-target,hx-current-utl")
		if r.Method == "OPTIONS" {
			// Respond to preflight request
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func groceryList(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Milk\nEggs\nBread\n")
}

func main() {
	http.HandleFunc("/hello", enableCors(hello))
	http.HandleFunc("/headers", enableCors(headers))
	http.HandleFunc("/groceryList", enableCors(groceryList))
	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}
