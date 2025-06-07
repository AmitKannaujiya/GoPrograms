package http

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HandleGreetFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "utf-8")
	w.Write([]byte("Hello Wellcom"))
}
func SetUpHttpRoute() {
	http.HandleFunc("/greet", HandleGreetFunc)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error loging ")
	}
}

func logMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("%s %s %s", r.Method, r.URL.Path, duration)
	})
}

func setupHttpRouteSample() {
	// path and handle func
	http.HandleFunc("/hello", hello)

	// listen to port
	fmt.Println("listening to port 8081")
	http.ListenAndServe(":8081", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

