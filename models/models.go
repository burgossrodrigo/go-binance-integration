package models

import "time"

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

type Trade struct {
    Type      string    `bson:"type"`
    Price     float64   `bson:"price"`
    Timestamp time.Time `bson:"timestamp"`
}

type Configs struct {
	MongoURI string
}