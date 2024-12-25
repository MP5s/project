package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calc_service/pkg/models"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           models.RequestBody
		expectedCode   int
		expectedResult string
	}{
		{
			name:           "Valid Expression",
			method:         http.MethodPost,
			body:           models.RequestBody{Expression: "2 + 2"},
			expectedCode:   http.StatusOK,
			expectedResult: `4.000000`,
		},
		{
			name:           "Invalid Expression",
			method:         http.MethodPost,
			body:           models.RequestBody{Expression: "2 + a"},
			expectedCode:   http.StatusUnprocessableEntity,
			expectedResult: `Expression is not valid`,
		},
		{
			name:           "Method Not Allowed",
			method:         http.MethodGet,
			body:           models.RequestBody{Expression: "2 + 2"},
			expectedCode:   http.StatusMethodNotAllowed,
			expectedResult: `Method not allowed`,
		},
	}

	for _, test := range tests {
		reqBody, _ := json.Marshal(test.body)
		req := httptest.NewRequest(test.method, "/calculate", bytes.NewBuffer(reqBody))
		w := httptest.NewRecorder()

		CalculateHandler(w, req)

		res := w.Result()

		if res.StatusCode != test.expectedCode {
			t.Errorf("%s: expected status %d, got %d", test.name, test.expectedCode, res.StatusCode)
		}

		var responseBody string
		if res.StatusCode == http.StatusOK {
			var response models.ResponseBody
			json.NewDecoder(res.Body).Decode(&response)
			responseBody = string(response.Result)
		} else {
			var errorResponse map[string]string
			json.NewDecoder(res.Body).Decode(&errorResponse)
			responseBody = errorResponse["error"]
		}

		if responseBody != test.expectedResult {
			t.Errorf("%s: expected body %q, got %q", test.name, test.expectedResult, responseBody)
		}
	}
}
