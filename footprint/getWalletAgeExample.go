package main

import (
	"fmt"
	"log"
	"net/url"

	"Web3DataAnalyse/consts"
	"Web3DataAnalyse/utils"
)

func main() {

	queryParamMap := map[string]string{
		"chain":          "Ethereum",
		"wallet_address": "0x46efbaedc92067e6d60e84ed6395099723252496",
	}

	baseURL := consts.ApiMap["getWalletAge"]
	query := url.Values{}
	for k, v := range queryParamMap {
		query.Add(k, v)
	}

	queryUrl := fmt.Sprintf("%s?%s", baseURL, query.Encode())

	header := map[string]string{
		"accept":  "application/json",
		"api-key": utils.Cfg.APIKey,
	}
	walletInfo, err := utils.GetUrlWithHeader(queryUrl, header)
	if err != nil {
		log.Println("getWalletAge query failed", err)
		return
	}

	// 示例输出结果
	//{
	//  "message": "success",
	//  "code": 0,
	//  "data": {
	//    "wallet_address": "0x46efbaedc92067e6d60e84ed6395099723252496",
	//    "age": 1024,
	//    "first_txn_time": "2021-06-22 16:31:21.000 UTC"
	//  }
	//}
	if nil != walletInfo {
		fmt.Println(walletInfo)
	}
}
