package main

// Simple static webserver
import (
    "log"
    "fmt"
    "net/http"
)
const PORT = "8080"


func handleCORS(h http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Access-Control-Allow-Origin", "*")
        h.ServeHTTP(w, r)
    }
}

func main() {
   fmt.Println("File server started and serving ./ at http://localhost:" + PORT) 
   log.Fatal(http.ListenAndServe(":"+PORT, handleCORS(http.FileServer(http.Dir("./")))))
}
