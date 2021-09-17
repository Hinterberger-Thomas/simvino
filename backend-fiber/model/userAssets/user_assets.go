package userAssets

type User_assets struct {
	User_assets_id uint32
	Fk_crypto_id   uint32
	Fk_user_id     uint32
	Amount         uint32
}

type User_assets_ins struct {
	Crypto_id  uint32
	Fk_user_id uint32
	Amount     uint32
}
