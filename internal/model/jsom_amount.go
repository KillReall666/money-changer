package model

type JSONRequest struct {
	Amount    float64
	Banknotes []float64
}

type JSONResponse struct {
	Exchanges [][]float64
}
