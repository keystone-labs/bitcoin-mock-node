# Bitcoin Mock Node

A mock Bitcoin RPC server that simulates Bitcoin node responses for testing purposes.

## Setup & Running

1. Install Go (version 1.22.3 or later)

2. Install dependencies:
```bash
go mod download
```

3. Run the server:
```bash
go run main.go
```

The server will output its URL, for example:
```
2024/11/27 16:22:55 Mock Bitcoin RPC server running at: http://127.0.0.1:52118
```

## Testing the Endpoints

You can test the RPC endpoints using curl. Replace `<PORT>` with the port number shown in the server output.

### Ping Test
```bash
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"MockServerHandler.Ping","params":[10],"id":1}' \
  http://127.0.0.1:<PORT>
```

### Get Best Block Hash
```bash
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"MockServerHandler.GetBestBlockHash","params":[],"id":1}' \
  http://127.0.0.1:<PORT>
```

### Get Block Count
```bash
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"MockServerHandler.GetBlockCount","params":[],"id":1}' \
  http://127.0.0.1:<PORT>
```

### Get Block Hash
```bash
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"MockServerHandler.GetBlockHash","params":[5],"id":1}' \
  http://127.0.0.1:<PORT>
```

## Running Tests

To run the test suite:
```bash
go test ./mockserver/...
```

## Data Sources

The mock server uses test data from:
- `data/mainnet_oldest_blocks.json`: Contains Bitcoin mainnet block headers
- `data/test.json`: Contains test UTXOs and transactions

## Note

The server will use a random available port each time it starts. Check the console output for the actual port number to use in your requests.