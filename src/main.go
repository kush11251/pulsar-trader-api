package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "pulsar-trader-api/src/controllers"
    "pulsar-trader-api/src/models"
    "pulsar-trader-api/src/utils"
)

func main() {
    http.HandleFunc("/api/metrics", controllers.GetMetrics)
    http.HandleFunc("/api/trades", controllers.GetTrades)
    fmt.Println("Server started on port 8080")
    http.ListenAndServe(":8080", nil)
}
