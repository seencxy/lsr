package utils

import (
	"math"
	"sms-activate/api"
)

type OrderBy int

const (
	Max OrderBy = iota
	Min
)

// QueryServiceFromMap query service from map
func QueryServiceFromMap(serviceMap map[string]*api.ServicePriceStruct, order OrderBy) ([]string, float64) {
	// sort by price
	priceMap := make(map[float64][]string)
	for serviceId, serviceInfo := range serviceMap {
		priceMap[serviceInfo.RetailPrice] = append(priceMap[serviceInfo.RetailPrice], serviceId)
	}

	// return the lowest or highest price service list according to the sorting
	var key float64
	switch order {
	case Max:
		maxKey := math.Inf(1)
		for k, _ := range priceMap {
			if k > maxKey {
				maxKey = k
			}
		}
		key = maxKey
	case Min:
		minKey := math.Inf(1)
		for k, _ := range priceMap {
			if k < minKey {
				minKey = k
			}
		}
		key = minKey
	}

	return priceMap[key], key
}
