package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/NucleonGodX/onchain-analytics/services"
)

type WalletHandler struct {
	analyticsService *services.AnalyticsService
}

func NewWalletHandler(analyticsService *services.AnalyticsService) *WalletHandler {
	return &WalletHandler{
		analyticsService: analyticsService,
	}
}

func (h *WalletHandler) GetAnalytics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	address := strings.TrimSpace(r.URL.Query().Get("address"))
	if address == "" {
		respondError(w, "Missing 'address' query parameter", http.StatusBadRequest)
		return
	}

	if !isValidEthAddress(address) {
		respondError(w, "Invalid Ethereum address format", http.StatusBadRequest)
		return
	}

	analytics, err := h.analyticsService.GetWalletAnalytics(address)
	if err != nil {
		log.Printf("Error getting analytics for %s: %v", address, err)
		respondError(w, "Failed to fetch analytics: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, analytics, http.StatusOK)
}

func isValidEthAddress(address string) bool {
	if !strings.HasPrefix(address, "0x") {
		return false
	}
	if len(address) != 42 {
		return false
	}
	return true
}

func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, map[string]string{"error": message}, status)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, map[string]string{
		"status": "healthy",
		"time":   time.Now().Format(time.RFC3339),
	}, http.StatusOK)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		respondError(w, "Not found", http.StatusNotFound)
		return
	}

	respondJSON(w, map[string]interface{}{
		"service": "On-Chain Analytics API",
		"version": "1.0.0",
		"endpoints": map[string]string{
			"analytics": "GET /analytics?address=<eth_address>",
			"health":    "GET /health",
		},
	}, http.StatusOK)
}