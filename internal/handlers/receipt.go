package handlers

import (
    "encoding/json"
    "net/http"
    "receipt-processor/internal/models"
    "receipt-processor/internal/services"
    "receipt-processor/internal/storage"

    "github.com/google/uuid"
    "github.com/gorilla/mux"
)

// ProcessReceipt handles the receipt processing request
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    id := uuid.New().String()
    storage.StoreReceipt(id, receipt)

    response := models.ReceiptResponse{ID: id}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// GetPoints retrieves the points awarded for a given receipt
func GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    receipt, exists := storage.GetReceipt(id)
    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    points := services.CalculatePoints(receipt)
    response := models.PointsResponse{Points: points}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
