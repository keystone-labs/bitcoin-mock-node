package main

import (
    "log"
    "path/filepath"
    "github.com/gonative-cc/btc-mock-node/mockserver"
)

func main() {
    // Create server with absolute path to data file
    dataPath := filepath.Join("data", "mainnet_oldest_blocks.json")
    server := mockserver.NewMockRPCServerWithPath(dataPath)
    defer server.Close()
    
    log.Printf("Mock Bitcoin RPC server running at: %s", server.URL)
    
    // Keep the server running
    select {}
}