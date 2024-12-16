package config

// DefaultValues is the default configuration
const DefaultValues = `
[log]
environment = "development" # "production" or "development"
level = "debug"
file = "/app/log/server.log"
outputs = ["stderr"]

[ethereum]
vizing-chain-id = 28516
networks = [
    {chain-id=28516,rpc="https://rpc-sepolia.vizing.com",entry-point="0xD4B426f113a858cD1521A8868d67D5407B3EdE5A",account-factory="0x75F82B744b3bD8b4A26AaECfB8157D996C7D1E95",sync-router="0xaAB80c905B5cd654132D3aC86b3922bA2F662d5F",is-sync=0,private_keys={path="/app/sender.keystore",password="password"}},
]

[db]
host = "127.0.0.1"
port = 3306
user = "root"
password = "sutai123456"
name = "spv"

[jsonrpc]
host = "0.0.0.0"
port = 8100

[pool]
max-ops = 2
flush-interval = 0

`
