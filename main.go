package main

import (
    utils "binance-integration/utils"
    binance "binance-integration/binance"
    "time"
)

func main() {
    ticker := time.NewTicker(6 * time.Hour)

    for {
        select {
        case <-ticker.C:
            ohlcData := binance.FetchOHLCData()
            resistance, support := utils.CalculateLevels(ohlcData)
            binance.ExecuteTrades(resistance, support)
        }
    }
}

// func executeTrades(resistance []float64, support []float64) {
//     client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
//     if err != nil {
//         log.Fatal(err)
//     }
//     ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//     err = client.Connect(ctx)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer client.Disconnect(ctx)

//     tradesCollection := client.Database("test").Collection("trades")

//     currentPrice := fetchCurrentPrice()

//     if len(resistance) < 1 || len(support) < 1 {
//         fmt.Println("No resistance or support levels found")
//         return
//     }

//     var tradeType string
//     if currentPrice > resistance[len(resistance)-1] {
//         fmt.Println("Sell Signal")
//         tradeType = "sell"
//     } else if currentPrice < support[len(support)-1] {
//         fmt.Println("Buy Signal")
//         tradeType = "buy"
//     } else {
//         fmt.Println("No Signal")
//         tradeType = "no signal"
//     }

//     trade := Trade{
//         Type:      tradeType,
//         Price:     currentPrice,
//         Timestamp: time.Now(),
//     }

//     _, err = tradesCollection.InsertOne(ctx, trade)
//     if err != nil {
//         log.Fatal(err)
//     }
// }

// func CalculateLevels(data []OHLC) ([]float64, []float64) {
// 	high := make([]float64, len(data))
// 	low := make([]float64, len(data))

// 	for i, ohlc := range data {
// 		high[i] = parseFloat64(ohlc.High)
// 		low[i] = parseFloat64(ohlc.Low)
// 	}

// 	period := 3 // Example period for the SMA
// 	if len(high) < period || len(low) < period {
// 		fmt.Println("Not enough data points for the SMA calculation")
// 		return
// 	}

// 	resistance := talib.Sma(high, period)
// 	support := talib.Sma(low, period)

// 	return resistance, support
// }

// func calculateProfit() {
//     client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
//     if err != nil {
//         log.Fatal(err)
//     }
//     ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//     err = client.Connect(ctx)
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer client.Disconnect(ctx)

//     tradesCollection := client.Database("test").Collection("trades")

//     cursor, err := tradesCollection.Find(ctx, bson.M{})
//     if err != nil {
//         log.Fatal(err)
//     }
//     var trades []Trade
//     if err = cursor.All(ctx, &trades); err != nil {
//         log.Fatal(err)
//     }

//     sort.Slice(trades, func(i, j int) bool {
//         return trades[i].Timestamp.Before(trades[j].Timestamp)
//     })

//     var totalProfit float64
//     for _, trade := range trades {
//         if trade.Type == "buy" {
//             totalProfit -= trade.Price
//         } else if trade.Type == "sell" {
//             totalProfit += trade.Price
//         }
//     }

//     fmt.Printf("Total profit: %f\n", totalProfit)
// }
