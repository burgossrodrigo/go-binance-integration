package utils

import (
	"fmt"
	"strconv"
	"github.com/markcheno/go-talib"
	models "binance-integration/models"
)

func ParseFloat64(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func CalculateLevels(data []models.OHLC) ([]float64, []float64) {
	high := make([]float64, len(data))
	low := make([]float64, len(data))

	for i, ohlc := range data {
		high[i] = ParseFloat64(ohlc.High)
		low[i] = ParseFloat64(ohlc.Low)
	}

	period := 3 // Example period for the SMA
	if len(high) < period || len(low) < period {
		fmt.Println("Not enough data points for the SMA calculation")
		return nil, nil
	}

	resistance := talib.Sma(high, period)
	support := talib.Sma(low, period)

	return resistance, support
}