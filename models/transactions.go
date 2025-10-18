package models

import "encoding/json"

type Transaction struct {
	Hash            string `json:"hash"`
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	Gas             string `json:"gas"`
	GasPrice        string `json:"gasPrice"`
	GasUsed         string `json:"gasUsed"`
	IsError         string `json:"isError"`
	TxReceiptStatus string `json:"txreceipt_status"`
}

type WalletAnalytics struct {
	Address        string  `json:"address"`
	TxCount        int     `json:"tx_count"`
	TotalETHSent   float64 `json:"total_eth_sent"`
	AvgGasUsed     int64   `json:"avg_gas_used"`
	HighestGasTx   string  `json:"highest_gas_tx"`
	SuccessfulTxns int     `json:"successful_txns"`
	FailedTxns     int     `json:"failed_txns"`
	TotalGasPaid   float64 `json:"total_gas_paid_eth"`
	AvgGasPrice    float64 `json:"avg_gas_price_gwei"`
	LastTxnTime    string  `json:"last_txn_time"`
}

type EtherscanResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type EtherscanTxResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Result  []Transaction `json:"result"`
}