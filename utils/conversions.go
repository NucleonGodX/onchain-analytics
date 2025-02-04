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