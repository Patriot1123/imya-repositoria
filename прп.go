package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

type Request struct {
    Expression string `json:"expression"`
}

type Response struct {
    Result string `json:"result,omitempty"`
    Error  string `json:"error,omitempty"`
}

func evaluateExpression(expr string) (float64, error) {
    // Простая проверка на допустимые символы (цифры и операции)
    if strings.ContainsAny(expr, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") {
        return 0, nil
    }
    // Здесь предположим, что выражение всегда корректное
    result, err := strconv.ParseFloat(expr, 64)
    return result, err
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    result, err := evaluateExpression(req.Expression)
    if err != nil {
        http.Error(w, `{"error":"Expression is not valid"}`, http.StatusUnprocessableEntity)
        return
    }

    response := Response{Result: strconv.FormatFloat(result, 'f', -1, 64)}
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/api/v1/calculate", calculateHandler)
    http.ListenAndServe(":8080", nil)
}