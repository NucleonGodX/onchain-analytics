package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/NucleonGodX/onchain-analytics/models"
)

type EtherscanService struct {
	apiKey     string
	httpClient *http.Client
	baseURL    string
}

func NewEtherscanService(apiKey string, timeout time.Duration) *EtherscanService {
	return &EtherscanService{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: "https://api.etherscan.io/v2/api",
	}
}

func (s *EtherscanService) FetchTransactions(address string) ([]models.Transaction, error) {
	url := fmt.Sprintf(
		"%s?chainid=1&module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=100&sort=desc&apikey=%s",
		s.baseURL, address, s.apiKey,
	)

	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transactions: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("etherscan API returned status %d", resp.StatusCode)
	}

	var genericResp struct {
		Status  string          `json:"status"`
		Message string          `json:"message"`
		Result  json.RawMessage `json:"result"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&genericResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if genericResp.Status != "1" {
		return nil, fmt.Errorf("etherscan API error: %s", genericResp.Message)
	}

	var txns []models.Transaction
	if err := json.Unmarshal(genericResp.Result, &txns); err != nil {
		var resultStr string
		if err2 := json.Unmarshal(genericResp.Result, &resultStr); err2 == nil {
			if resultStr == "No transactions found" {
				return []models.Transaction{}, nil
			}
			return nil, fmt.Errorf("etherscan returned: %s", resultStr)
		}
		return nil, fmt.Errorf("failed to parse transactions: %w", err)
	}

	return txns, nil
}