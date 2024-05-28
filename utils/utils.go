package utils

import (
	"fmt"
	"strconv"
	"github.com/markcheno/go-talib"
	models "binance-integration/models"
	"sort"
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

func IdentifyLevels(ohlcData []models.OHLC) (resistance []float64, support []float64) {
	var highs, lows []float64

	for _, ohlc := range ohlcData {
		highs = append(highs, ParseFloat64(ohlc.High))
		lows = append(lows, ParseFloat64(ohlc.Low))
	}

	sort.Float64s(highs)
	sort.Float64s(lows)

	// Identify top 3 resistance levels (highest highs)
	for i := len(highs) - 3; i < len(highs); i++ {
		resistance = append(resistance, highs[i])
	}

	// Identify bottom 3 support levels (lowest lows)
	for i := 0; i < 3; i++ {
		support = append(support, lows[i])
	}

	return resistance, support
}