package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type OHLC struct {
	OpenTime                float64
	Open                    string
	High                    string
	Low                     string
	Close                   string
	Volume                  string
	CloseTime               float64
	BaseAssetVolume         string
	NumberOfTrades          float64
	TakerBuyVolume          string
	TakerBuyBaseAssetVolume string
	Ignore                  string
}

func main() {
	url := "https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=1m&limit=32"

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		panic("Error doing request")
	}

	bodyBytes := resp.Body()

	var responses [][]interface{}
	err := json.Unmarshal(bodyBytes, &responses)
	if err != nil {
		fmt.Print("Error unmarshaling response", err)
	}

	var formattedResponses []OHLC
	for _, response := range responses {
		formattedResponse := OHLC{
			OpenTime:                response[0].(float64),
			Open:                    response[1].(string),
			High:                    response[2].(string),
			Low:                     response[3].(string),
			Close:                   response[4].(string),
			Volume:                  response[5].(string),
			CloseTime:               response[6].(float64),
			BaseAssetVolume:         response[7].(string),
			NumberOfTrades:          response[8].(float64),
			TakerBuyVolume:          response[9].(string),
			TakerBuyBaseAssetVolume: response[10].(string),
			Ignore:                  response[11].(string),
		}
		formattedResponses = append(formattedResponses, formattedResponse)
	}

	fmt.Print(formattedResponses)
}
