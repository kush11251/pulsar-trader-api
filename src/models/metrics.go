package models

import (
    "encoding/json"
)

type Metric struct {
    Name  string `json:"name"`
    Value int    `json:"value"`
}

func GetMetrics() []Metric {
    metrics := []Metric{{Name: "metric1", Value: 10}, {Name: "metric2", Value: 20}}
    return metrics
}
