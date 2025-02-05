package utils

import (
	"strconv"
	"strings"
)

const (
	WeiPerEth  = 1e18
	WeiPerGwei = 1e9
)

func WeiToEth(wei string) float64 {
	value, err := strconv.ParseFloat(wei, 64)
	if err != nil {
		return 0
	}
	return value / WeiPerEth
}

func WeiToGwei(wei string) float64 {
	value, err := strconv.ParseFloat(wei, 64)
	if err != nil {
		return 0
	}
	return value / WeiPerGwei
}

func CalculateGasFee(gasUsed, gasPrice string) float64 {
	used, err1 := strconv.ParseFloat(gasUsed, 64)
	price, err2 := strconv.ParseFloat(gasPrice, 64)
	
	if err1 != nil || err2 != nil {
		return 0
	}
	
	return (used * price) / WeiPerEth
}

func ParseInt64(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func NormalizeAddress(addr string) string {
	return strings.ToLower(strings.TrimSpace(addr))
}