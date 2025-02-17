package services

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
    "receipt-processor/internal/models"
)

// CalculatePoints applies the business logic to determine receipt points
func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // 1 point per alphanumeric character in retailer name
    re := regexp.MustCompile(`[a-zA-Z0-9]`)
    points += len(re.FindAllString(receipt.Retailer, -1))

    // Convert total to float
    total, _ := strconv.ParseFloat(receipt.Total, 64)

    // 50 points if total is a round dollar amount
    if total == float64(int(total)) {
        points += 50
    }

    // 25 points if total is a multiple of 0.25
    if math.Mod(total, 0.25) == 0 {
        points += 25
    }

    // 5 points for every two items
    points += (len(receipt.Items) / 2) * 5

    // Points for item descriptions
    for _, item := range receipt.Items {
        desc := strings.TrimSpace(item.ShortDescription)
        if len(desc)%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(price * 0.2))
        }
    }

    // Parse purchase date
    date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
    if date.Day()%2 != 0 {
        points += 6
    }

    // Parse purchase time
    purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
    if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
        points += 10
    }

    return points
}
