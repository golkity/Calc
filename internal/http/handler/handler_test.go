package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalcHandler_Success(t *testing.T) {
	testCases := []struct {
		reqBody        string
		expectedResult float64
	}{
		{`{"expression":"1+1"}`, 2},
		{`{"expression":"2*3+(2+2)"}`, 10},
		{`{"expression":"10/2"}`, 5},
		{`{"expression":"5-3"}`, 2},
		{`{"expression":"100*2-50"}`, 150},
		{`{"expression":"(1+2)*3"}`, 9},
	}

	for _, testCase := range testCases {
		t.Run(testCase.reqBody, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(testCase.reqBody)))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			CalcHandler(resp, req)

			if resp.Code != http.StatusOK {
				t.Errorf("expected status 200, got %d", resp.Code)
			}

			var result map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if result["result"] != testCase.expectedResult {
				t.Errorf("expected result %v, got %v", testCase.expectedResult, result["result"])
			}
		})
	}
}

func TestCalcHandler_Errors(t *testing.T) {
	testCases := []struct {
		reqBody       string
		expectedError string
		expectedCode  int
	}{
		{`{"expression":"1+"}`, "Invalid expression: invalid syntax", http.StatusUnprocessableEntity},
		{`{"expression":"2*+(2+2)"}`, "Invalid expression: invalid syntax", http.StatusUnprocessableEntity},
		{`{"expression":""}`, "Expression cannot be empty", http.StatusBadRequest},
		{`invalid json`, "Invalid JSON format", http.StatusBadRequest},
	}

	for _, testCase := range testCases {
		t.Run(testCase.reqBody, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(testCase.reqBody)))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			CalcHandler(resp, req)

			if resp.Code != testCase.expectedCode {
				t.Errorf("expected status %d, got %d", testCase.expectedCode, resp.Code)
			}

			var result map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if result["error"] != testCase.expectedError {
				t.Errorf("expected error %q, got %q", testCase.expectedError, result["error"])
			}
		})
	}
}

func TestCalcHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/calculate", nil)
	resp := httptest.NewRecorder()

	CalcHandler(resp, req)

	if resp.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status 405, got %d", resp.Code)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if result["error"] != "Only POST method is allowed" {
		t.Errorf("expected error %q, got %q", "Only POST method is allowed", result["error"])
	}
}
