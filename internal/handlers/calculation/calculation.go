package calculation

import (
	"encoding/json"
	"io"
	"net/http"

	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/model"
)

type Calculater interface {
	Calculate(data model.JSONRequest) ([][]float64, error)
}

type CalculationHandler struct {
	Calculater Calculater
	log        *logger.Logger
}

func NewCalculationHandler(calc Calculater, log *logger.Logger) *CalculationHandler {
	return &CalculationHandler{
		Calculater: calc,
		log:        log,
	}
}

func (h *CalculationHandler) GetCalculation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST requests support!", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusInternalServerError)
		return
	}

	amountAndBanknotes := model.JSONRequest{}

	err = json.Unmarshal(body, &amountAndBanknotes)
	if err != nil {
		http.Error(w, "could not parse request body", http.StatusBadRequest)
		return
	}

	exchanges := model.JSONResponse{}

	exchanges.Exchanges, err = h.Calculater.Calculate(amountAndBanknotes)
	if err != nil {
		http.Error(w, "could not calculate exchanges: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(exchanges)
	if err != nil {
		http.Error(w, "could not encode request body", http.StatusInternalServerError)
		return
	}

}
