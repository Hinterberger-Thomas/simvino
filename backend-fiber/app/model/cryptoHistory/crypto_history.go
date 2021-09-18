package cryptoHistory

import "time"

type CryptoHistroy struct {
	Crypto_history_id uint64
	Price             float64
	Fk_crypto_id      uint32
	Date              time.Time
}

type CryptoHistroyIns struct {
	Asset_id string  `bson:"asset_id"`
	Price    float64 `bson:"price"`
	Time     int64   `bson:"time"`
}
