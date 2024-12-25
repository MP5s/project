package handlers

import (
	"encoding/json"
	"net/http"

	"calc_service/internal/application"
	"calc_service/pkg/calculator"
	"calc_service/pkg/models"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var reqBody models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
		return
	}

	if !calculator.IsValidExpression(reqBody.Expression) {
		http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := application.ProcessExpression(reqBody.Expression)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	response := models.ResponseBody{Result: result}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
