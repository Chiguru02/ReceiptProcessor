package main

import (
    "fmt"
    "log"
    "net/http"
    "receipt-processor/internal/handlers"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

    fmt.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
