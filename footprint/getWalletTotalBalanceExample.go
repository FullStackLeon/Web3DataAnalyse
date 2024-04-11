package main

import (
	"fmt"
	"log"
	"net/url"

	"Web3DataAnalyse/consts"
	"Web3DataAnalyse/utils"
)

func main() {
	baseURL := consts.ApiMap["getWalletTotalBalance"]
	queryParamMap := map[string]string{
		"chain":          "Ethereum",
		"wallet_address": "0x46efbaedc92067e6d60e84ed6395099723252496",
	}
	query := url.Values{}
	for k, v := range queryParamMap {
		query.Add(k, v)
	}
	queryUrl := fmt.Sprintf("%s?%s", baseURL, query.Encode())

	header := map[string]string{
		"accept":  "application/json",
		"api-key": utils.Cfg.APIKey,
	}
	walletTotalBalanceInfo, err := utils.GetUrlWithHeader(queryUrl, header)
	if err != nil {
		log.Println("getWalletAge query failed", err)
		return
	}
	if nil != walletTotalBalanceInfo {
		fmt.Println(walletTotalBalanceInfo)
	}
	// 示例输出结果
	//{
	//  "message": "success",
	//  "code": 0,
	//  "data": {
	//    "value": 23099.92299537051
	//  }
	//}
}
