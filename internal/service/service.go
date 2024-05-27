package service

import (
	"fmt"
	"sort"

	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/config"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/logger"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/model"
)

type Service struct {
	cfg *config.Config
	log *logger.Logger
}

func New(cfg *config.Config, log *logger.Logger) *Service {
	return &Service{
		cfg: cfg,
		log: log,
	}
}

func (s *Service) Calculate(data model.JSONRequest) ([][]float64, error) {
	if int(data.Amount)%50 > 0 {
		return nil, fmt.Errorf("amount must be a multiple of 50")
	} else if len(data.Banknotes) <= 0 {
		return nil, fmt.Errorf("banknotes cannot be empty")
	} else if data.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero")
	}

	target := int(data.Amount * 100)
	banknoteCombinations := make([][][]float64, target+1)
	banknoteCombinations[0] = append(banknoteCombinations[0], []float64{})

	for _, bn := range data.Banknotes {
		note := int(bn * 100)
		for i := note; i <= target; i++ {
			for _, comb := range banknoteCombinations[i-note] {
				newComb := append(comb, bn)
				sort.Float64s(newComb) //
				banknoteCombinations[i] = append(banknoteCombinations[i], newComb)
			}
		}
	}

	return banknoteCombinations[target], nil
}
