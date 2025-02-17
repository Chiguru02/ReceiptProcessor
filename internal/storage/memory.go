package storage

import (
    "receipt-processor/internal/models"
    "sync"
)

var (
    receiptStore = make(map[string]models.Receipt)
    storeLock    = sync.Mutex{}
)

// StoreReceipt saves a receipt with an ID
func StoreReceipt(id string, receipt models.Receipt) {
    storeLock.Lock()
    receiptStore[id] = receipt
    storeLock.Unlock()
}

// GetReceipt retrieves a receipt by ID
func GetReceipt(id string) (models.Receipt, bool) {
    storeLock.Lock()
    receipt, exists := receiptStore[id]
    storeLock.Unlock()
    return receipt, exists
}
