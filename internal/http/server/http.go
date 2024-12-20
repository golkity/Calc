package server

import (
	"fmt"
	"github.com/golkity/Calc/internal/http/handler"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type SuccessResponse struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func RegRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/calculate", handler.CalcHandler)
}

func formatResult(result float64) string {
	return string([]byte(fmt.Sprintf("%.6f", result)))
}
