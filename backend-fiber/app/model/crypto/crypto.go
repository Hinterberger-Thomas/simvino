package crypto

type Crypto struct {
	Crypto_id     uint32
	Asset_id      string
	Asset_name    string
	IS_CRYPTO     int8
	Current_price float64
}

type CryptoIns struct {
	Asset_id      string
	Asset_name    string
	Is_crypto     int8
	Current_price float64
}

type CryptoUpd struct {
	Asset_id      string
	Current_price float64
}

type CryptoJson struct {
	AssetID   string  `json:"asset_id"`
	Name      string  `json:"name"`
	PriceUsd  float64 `json:"price_usd"`
	Is_crypto int8    `json:"type_is_crypto"`
}
