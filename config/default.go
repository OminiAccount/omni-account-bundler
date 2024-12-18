package config

// DefaultValues is the default configuration
const DefaultValues = `
[log]
environment = "development" # "production" or "development"
level = "debug"
file = "/app/log/server.log"
outputs = ["stderr"]

[ethereum]
[ethereum.vizing]
chain-id = 28516
rpc = "https://rpc-sepolia.vizing.com"
entry-point = "0x03064C8cDB2dd14Edcd225204BB5f45931D26783"
account-factory = "0xD64986da346370b4B24Fe23e25b57A5676782157"
sync-router = "0xF126c02D118A1d2788B6eF3608E858aCF5A6BF0E"
is-sync = 1
gen-block = 0
    [ethereum.vizing.private_keys]
    path = "./bundler.keystore"
    password = "password"

[jsonrpc]
host = "0.0.0.0"
port = 8100

[pool]
max-ops = 2
flush-interval = 0

[state]
max-batches = 10
`
