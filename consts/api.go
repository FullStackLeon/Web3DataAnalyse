package consts

var (
	FootPrintBaseUrl = "https://api.footprint.network/api"
	ApiMap           = map[string]string{
		"getWalletAge":          "/v3/address/getWalletAge",
		"getWalletTotalBalance": "/v3/address/getWalletTotalBalance",
		"getProtocolMetadata":   "/v2/protocol/info",
	}
)
