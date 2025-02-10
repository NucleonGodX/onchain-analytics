package services

import (
	"sync"
	"time"

	"github.com/yourusername/onchain-analytics/models"
	"github.com/yourusername/onchain-analytics/utils"
)

type AnalyticsService struct {
	etherscan *EtherscanService
	cache     map[string]*CachedAnalytics
	cacheTTL  time.Duration
	mu        sync.RWMutex
}

type CachedAnalytics struct {
	Data      *models.WalletAnalytics
	ExpiresAt time.Time
}

func NewAnalyticsService(etherscan *EtherscanService, cacheTTL time.Duration) *AnalyticsService {
	return &AnalyticsService{
		etherscan: etherscan,
		cache:     make(map[string]*CachedAnalytics),
		cacheTTL:  cacheTTL,
	}
}

func (s *AnalyticsService) GetWalletAnalytics(address string) (*models.WalletAnalytics, error) {
	normalizedAddr := utils.NormalizeAddress(address)

	s.mu.RLock()
	cached, exists := s.cache[normalizedAddr]
	s.mu.RUnlock()

	if exists && time.Now().Before(cached.ExpiresAt) {
		return cached.Data, nil
	}

	txns, err := s.etherscan.FetchTransactions(address)
	if err != nil {
		return nil, err
	}

	analytics := s.computeAnalytics(normalizedAddr, txns)

	s.mu.Lock()
	s.cache[normalizedAddr] = &CachedAnalytics{
		Data:      analytics,
		ExpiresAt: time.Now().Add(s.cacheTTL),
	}
	s.mu.Unlock()

	return analytics, nil
}

func (s *AnalyticsService) computeAnalytics(address string, txns []models.Transaction) *models.WalletAnalytics {
	analytics := &models.WalletAnalytics{
		Address: address,
		TxCount: len(txns),
	}

	if len(txns) == 0 {
		return analytics
	}

	var (
		totalGasUsed   int64
		totalGasPrice  float64
		highestGasFee  float64
		highestGasTxHash string
	)

	for _, tx := range txns {
		if tx.IsError == "0" && tx.TxReceiptStatus == "1" {
			analytics.SuccessfulTxns++
		} else {
			analytics.FailedTxns++
		}

		if utils.NormalizeAddress(tx.From) == address {
			analytics.TotalETHSent += utils.WeiToEth(tx.Value)
		}

		if utils.NormalizeAddress(tx.From) == address {
			gasFee := utils.CalculateGasFee(tx.GasUsed, tx.GasPrice)
			analytics.TotalGasPaid += gasFee

			if gasFee > highestGasFee {
				highestGasFee = gasFee
				highestGasTxHash = tx.Hash
			}

			totalGasUsed += utils.ParseInt64(tx.GasUsed)
			totalGasPrice += utils.WeiToGwei(tx.GasPrice)
		}

		if analytics.LastTxnTime == "" {
			timestamp := utils.ParseInt64(tx.TimeStamp)
			analytics.LastTxnTime = time.Unix(timestamp, 0).Format(time.RFC3339)
		}
	}

	if analytics.TxCount > 0 {
		analytics.AvgGasUsed = totalGasUsed / int64(analytics.TxCount)
		analytics.AvgGasPrice = totalGasPrice / float64(analytics.TxCount)
	}

	analytics.HighestGasTx = highestGasTxHash

	return analytics
}

func (s *AnalyticsService) ClearCache() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache = make(map[string]*CachedAnalytics)
}