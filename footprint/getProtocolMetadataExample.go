package main

import (
	"fmt"
	"log"
	"net/url"

	"Web3DataAnalyse/consts"
	"Web3DataAnalyse/utils"
)

func main() {
	baseURL := consts.FootPrintBaseUrl + consts.ApiMap["getProtocolMetadata"]
	queryParamMap := map[string]string{
		"chain":         "Ethereum",
		"protocol_slug": "the-sandbox",
		"limit":         "1",
		"offset":        "10",
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
	protocolMetadataInfo, err := utils.GetUrlWithHeader(queryUrl, header)
	if err != nil {
		log.Println("getProtocolMetadata query failed", err)
		return
	}

	if nil != protocolMetadataInfo {
		fmt.Println(protocolMetadataInfo)
	}
	// 示例输出结果
	//{
	//  "message": "success",
	//  "code": 0,
	//  "data": [
	//    {
	//      "chain": "Ethereum",
	//      "protocol_slug": "the-sandbox",
	//      "protocol_type": "GameFi",
	//      "protocol_sub_type": null,
	//      "protocol_name": "The Sandbox",
	//      "logo": "https://footprint-imgs.oss-us-east-1.aliyuncs.com/logo_images/the-sandbox.jpg",
	//      "categories": [
	//        "Minigame",
	//        "Open-World",
	//        "Virtual-World"
	//      ],
	//      "discord": "https://discordapp.com/invite/vAe4zvY",
	//      "github": null,
	//      "twitter": "TheSandboxGame",
	//      "telegram": null,
	//      "web_url": "https://www.sandbox.game/en/",
	//      "description": "The Sandbox is a virtual gaming world where players can build, own, and monetize their gaming experiences.",
	//      "updated_at": "2024-03-18 12:17:34.000 UTC",
	//      "created_at": "2023-05-08 06:50:48.000 UTC"
	//    }
	//  ]
	//}
}
