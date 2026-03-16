package controllers

import (
    "encoding/json"
    "net/http"
    "pulsar-trader-api/src/models"
    "pulsar-trader-api/src/utils"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
    metrics := models.GetMetrics()
    json.NewEncoder(w).Encode(metrics)
}
