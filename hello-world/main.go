package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Read the HTML file
    html, err := ioutil.ReadFile("index.html")
    if err != nil {
        http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
        return
    }

    // Set the content type to HTML and write the HTML content
    w.Header().Set("Content-Type", "text/html")
    w.Write(html)
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Server is listening on port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
