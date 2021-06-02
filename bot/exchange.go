package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StatusType uint8

const (
	Running StatusType = 0
	Offline StatusType = 1
)

type Exchange interface {
	Status() StatusType
	Assets() []string
}

type BinanceExchange struct {
	Endpoint  string
	apiKey    string
	apiSecret string
}

func (b BinanceExchange) Status() StatusType {
	resp, err := http.Get(fmt.Sprintf("%s/%s", b.Endpoint, "api/v3/ping"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("Got response code ", resp.StatusCode)
		return Offline
	}

	return Running
}

func (b BinanceExchange) Assets() []string {
	type responseData struct {
		Symbols []struct {
			Symbol    string `json:"symbol"`
			Status    string `json:"status"`
			BaseAsset string `json:"baseAsset"`
		}
	}

	resp, err := http.Get(fmt.Sprintf("%s/%s", b.Endpoint, "api/v3/exchangeInfo"))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := responseData{}
	json.Unmarshal(body, &data)

	var assets []string
	for i := 0; i < len(data.Symbols); i++ {
		assets = append(assets, data.Symbols[i].BaseAsset)
	}
	return assets
}
