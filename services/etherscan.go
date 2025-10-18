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
		baseURL: "https://api.etherscan.io/api",
	}
}

func (s *EtherscanService) FetchTransactions(address string) ([]models.Transaction, error) {
	url := fmt.Sprintf(
		"%s?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&page=1&offset=100&sort=desc&apikey=%s",
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

	var ethResp models.EtherscanResponse
	if err := json.NewDecoder(resp.Body).Decode(&ethResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if ethResp.Status != "1" {
		return nil, fmt.Errorf("etherscan API error: %s", ethResp.Message)
	}

	return ethResp.Result, nil
}