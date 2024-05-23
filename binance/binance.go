package binance

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	models "binance-integration/models"
	utils "binance-integration/utils"
	config "binance-integration/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
	"time"
)


func FetchOHLCData() []models.OHLC {
	url := "https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=6h&limit=32"

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		panic("Error doing request: " + err.Error())
	}

	bodyBytes := resp.Body()

	var responses [][]interface{}
	err := json.Unmarshal(bodyBytes, &responses)
	if err != nil {
		fmt.Print("Error unmarshaling response: ", err)
		return nil
	}

	// Check if responses have at least one element
	if len(responses) == 0 {
		fmt.Println("No data received in response")
		return nil
	}

	var formattedResponses []models.OHLC
	for _, response := range responses {
		formattedResponse := models.OHLC{
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

	open := make([]float64, len(formattedResponses))
	high := make([]float64, len(formattedResponses))
	low := make([]float64, len(formattedResponses))
	close := make([]float64, len(formattedResponses))

	for i, ohlc := range formattedResponses {
		open[i] = utils.ParseFloat64(ohlc.Open)
		high[i] = utils.ParseFloat64(ohlc.High)
		low[i] = utils.ParseFloat64(ohlc.Low)
		close[i] = utils.ParseFloat64(ohlc.Close)
	}

	return formattedResponses
}

func FetchCurrentPrice() float64 {
	url := "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT"

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}
	if err := client.Do(req, resp); err != nil {
		panic("Error doing request: " + err.Error())
	}

	bodyBytes := resp.Body()

	var priceData struct {
		Price string `json:"price"`
	}
	err := json.Unmarshal(bodyBytes, &priceData)
	if err != nil {
		fmt.Print("Error unmarshaling response: ", err)
		return 0
	}

	return utils.ParseFloat64(priceData.Price)
}

func ExecuteTrades(resistance []float64, support []float64) {
    cfg := config.LoadEnv()
	
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoURI))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    tradesCollection := client.Database("test").Collection("trades")

    currentPrice := FetchCurrentPrice()

    if len(resistance) < 1 || len(support) < 1 {
        fmt.Println("No resistance or support levels found")
        return
    }

    var tradeType string
    if currentPrice > resistance[len(resistance)-1] {
        fmt.Println("Sell Signal")
        tradeType = "sell"
    } else if currentPrice < support[len(support)-1] {
        fmt.Println("Buy Signal")
        tradeType = "buy"
    } else {
        fmt.Println("No Signal")
        tradeType = "no signal"
    }

    trade := models.Trade{
        Type:      tradeType,
        Price:     currentPrice,
        Timestamp: time.Now(),
    }

    _, err = tradesCollection.InsertOne(ctx, trade)
    if err != nil {
        log.Fatal(err)
    }
}
