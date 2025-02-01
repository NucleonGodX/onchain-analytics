package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/yourusername/onchain-analytics/handlers"
	"github.com/yourusername/onchain-analytics/services"
)

func main() {
	apiKey := os.Getenv("ETHERSCAN_API_KEY")
	if apiKey == "" {
		log.Fatal("ETHERSCAN_API_KEY environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	etherscanService := services.NewEtherscanService(apiKey, 10*time.Second)
	analyticsService := services.NewAnalyticsService(etherscanService, 5*time.Minute)

	walletHandler := handlers.NewWalletHandler(analyticsService)

	mux := http.NewServeMux()
	mux.HandleFunc("/analytics", walletHandler.GetAnalytics)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/", handlers.RootHandler)

	handler := loggingMiddleware(mux)

	log.Printf("🚀 Server starting on port %s", port)
	log.Printf("📊 Endpoints:")
	log.Printf("  GET /analytics?address=<eth_address>")
	log.Printf("  GET /health")

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}