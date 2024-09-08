# omni-account-bundler
This Go-based implementation serves as a basic EIP-4337 bundler, responsible for bundling user operations and interacting with L1 and ZKP systems. It handles the aggregation and submission of transactions while ensuring secure and efficient communication with the underlying layers.

## Getting Started

1. Clone this repo and start off in the `omni-account-bundler` directory.

    ```bash
    git clone https://github.com/OminiAccount/omni-account-bundler.git
    cd omni-account-bundler
    ```

2. Configure environment file

   Rename `example.toml` to `xx.toml`. Such as `develop.toml`.

3. Start the processor

    ```bash
    cd cmd
    go build
    
    #develop 
    nohup ./cmd processor > cmd.log 2>&1 &
    #release
    nohup ./cmd processor --c ../../develop.toml > cmd.log 2>&1 &
    ```