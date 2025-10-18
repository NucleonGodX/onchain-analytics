# On-Chain Analytics API

A lightweight Go service that analyzes Ethereum wallet transactions and gas metrics using the Etherscan API.

## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- Etherscan API key (free at https://etherscan.io/apis)

### Setup

```bash
# 1. Create project structure
mkdir -p go-chain-analytics/{handlers,services,models,utils}
cd go-chain-analytics

# 2. Initialize Go module
go mod init github.com/yourusername/onchain-analytics

# 3. Set environment variables
export ETHERSCAN_API_KEY=your_api_key_here
export PORT=8080

# 4. Run the server
go run .
```

### Test Locally

```bash
# Health check
curl http://localhost:8080/health

# Analyze a wallet
curl "http://localhost:8080/analytics?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
```

## 📡 API Endpoints

### GET `/analytics?address=<wallet_address>`

Returns comprehensive wallet analytics including transaction count, ETH sent, and gas metrics.

**Example Request:**
```bash
curl "https://onchain-analytics.onrender.com/analytics?address=0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
```

**Example Response:**
```json
{
  "address": "0xd8da6bf26964af9d7eed9e03e53415d37aa96045",
  "tx_count": 100,
  "total_eth_sent": 50,
  "avg_gas_used": 16743,
  "highest_gas_tx": "0x87a41deccb73004da4615718ecce5f44f802f8c11125ff68c94a83f03b5583b6",
  "successful_txns": 93,
  "failed_txns": 7,
  "total_gas_paid_eth": 0.000314271843014466,
  "avg_gas_price_gwei": 0.02736750148,
  "last_txn_time": "2025-10-18T12:11:59Z"
}
```

## 🐳 Live Deployment

**Service:** Render  
**URL:** https://onchain-analytics.onrender.com

The API is live and ready to use. Simply pass any Ethereum address to analyze its on-chain activity.

## 📊 Response Fields

- `address` - Normalized Ethereum wallet address
- `tx_count` - Total number of transactions
- `total_eth_sent` - Total ETH amount sent from wallet
- `avg_gas_used` - Average gas consumed per transaction
- `highest_gas_tx` - Transaction hash with highest gas consumption
- `successful_txns` - Count of successful transactions
- `failed_txns` - Count of failed transactions
- `total_gas_paid_eth` - Total ETH spent on gas fees
- `avg_gas_price_gwei` - Average gas price in Gwei
- `last_txn_time` - Timestamp of most recent transaction