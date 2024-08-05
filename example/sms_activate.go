package main

import (
	"github.com/seencxy/lsr/sms_activate/api"
	"github.com/seencxy/lsr/sms_activate/utils"
	"log"
	"net/http"
)

func main() {
	const apiKey = ""

	client := api.NewClient(http.Client{})

	number, err := client.GetPhoneNumber(apiKey, "6", "any", "dr")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(number)

	service, err := client.QueryTopCountriesByService(apiKey, "dr", false)
	if err != nil {
		log.Println(err)
		return
	}
	// 需要对比这个国家里面的 retail_price 就是价格
	countryList, price := utils.QueryServiceFromMap(service, utils.Min)
	log.Println(countryList, price)
	if price != 10 {
		log.Println("The price is 10.")
		return
	}
	for _, country := range countryList {
		if country == "16" {
			number, err := client.GetPhoneNumber(apiKey, country, "any", "dr")
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(number)
		}
	}
}
