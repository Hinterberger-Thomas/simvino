package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Hinterberger-Thomas/simvino/model/crypto"
	"github.com/Hinterberger-Thomas/simvino/model/cryptoHistory"
	resty "github.com/go-resty/resty/v2"
)

func GetCoinsValue() {
	for {
		// resp, err := resty.R().
		// 	SetHeader("X-CoinAPI-Key", "73034021-THIS-IS-SAMPLE-KEY").
		// 	Get("https://rest.coinapi.io/v1/assets")

		resp, err := resty.New().R().
			SetHeader("X-CoinAPI-Key", "1017114D-4B5D-4763-BEB4-2F60F9072DD0").
			Get("https://rest.coinapi.io/v1/assets")
		//var cryptos crypto.CryptoJson
		if err != nil {
			fmt.Println(err)
		}
		var cryptos []crypto.CryptoJson

		err = json.Unmarshal(resp.Body(), &cryptos)
		if err != nil {
			fmt.Println(err)
		}

		var hisInter []interface{}
		cryptoHistory.Get(time.Now().Add(time.Hour*-24), time.Now(), "btc")
		crypto.InsertNewCrypto(cryptos)

		for i := 0; i < len(cryptos); i++ {
			hisInter = append(hisInter, cryptoHistory.CryptoHistroyIns{Asset_id: cryptos[i].AssetID, Price: cryptos[i].PriceUsd, Time: time.Now().Unix()})
			crypto.UpdateCryptoByAssetId(crypto.CryptoUpd{Asset_id: cryptos[i].AssetID, Current_price: cryptos[i].PriceUsd})
		}

		cryptoHistory.InsertHis(hisInter)

		// 	err := crypto.InsertNewCrypto(crypto.CryptoIns{Asset_id: cryptos[i].AssetID, Asset_name: cryptos[i].Name, Current_price: cryptos[i].PriceUsd, Is_crypto: cryptos[i].Is_crypto})
		// 	hisInter = append(hisInter, cryptoHistory.CryptoHistroyIns{Asset_id: cryptos[i].AssetID, Price: cryptos[i].PriceUsd, Time: time.Now()})
		// 	if err != nil {
		// 		err = crypto.UpdateCryptoByAssetId(crypto.CryptoUpd{Asset_id: cryptos[i].AssetID, Current_price: cryptos[i].PriceUsd})
		// 		if err != nil {
		// 			fmt.Println(err)
		// 		}
		// 	}
		// 	err = cryptoHistory.InsertCrypto(cryptoHistory.CryptoHistroyIns{Price: cryptos[i].PriceUsd, Asset_id: cryptos[i].AssetID})
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		time.Sleep(25 * time.Minute)
		// }

	}
}
