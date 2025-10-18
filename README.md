## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Etherscan API key (free at https://etherscan.io/apis)

### Setup
```bash
# 1. Create project structure
mkdir -p go-chain-analytics/{handlers,services,models,utils}
cd go-chain-analytics

# 2. Copy all files into their directories

# 3. Initialize module
go mod init github.com/yourusername/onchain-analytics

# 4. Set environment variables
export ETHERSCAN_API_KEY=your_api_key_here
export PORT=8080

# 5. Run
go run .
```

### Test
```bash
# Health check
curl http://localhost:8080/health

# Analyze wallet
curl "http://localhost:8080/analytics?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
```

## 📡 API Endpoints

### GET /analytics?address=<eth_address>

Returns wallet analytics including transaction count, ETH sent, gas metrics.

**Example Response:**
```json
{
  "address": "0xabc123...",
  "tx_count": 25,
  "total_eth_sent": 2.43,
  "avg_gas_used": 42187,
  "highest_gas_tx": "0x9fa...",
  "successful_txns": 23,
  "failed_txns": 2,
  "total_gas_paid_eth": 0.0234,
  "avg_gas_price_gwei": 28.5,
  "last_txn_time": "2025-01-15T10:30:00Z"
}
```

## 🐳 Deployment

### Render
Deployed on render

---
