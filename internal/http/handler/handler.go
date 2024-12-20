package handler

import (
	"Calc/pkg/calc"
	"Calc/pkg/logger"
	"encoding/json"
	"net/http"
)

var (
	lg *logger.Logger
)

type Request struct {
	Expression string `json:"Expression"`
}

type Response struct {
	Result *float64 `json:"result"`
	Error  string   `json:"error"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	lg.Info("Received request:", r.Method, r.URL.Path)

	if r.Method != http.MethodPost {
		lg.Info("Invalid method used:", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Only POST method is allowed",
		})
		return
	}

	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		lg.Info("Failed to decode request body:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	lg.Debug("Expression received:", req.Expression)

	if req.Expression == "" {
		lg.Info("Empty expression received")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Expression cannot be empty",
		})
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		lg.Info("Calculation error for expression:", req.Expression, "Error:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid expression: " + err.Error(),
		})
		return
	}

	lg.Info("Calculation successful for expression:", req.Expression, "Result:", result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": result,
	})
}
