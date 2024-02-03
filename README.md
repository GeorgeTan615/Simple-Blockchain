# Simple-Blockchain
A blockchain project built to gain understanding on the fundamental workings of blockchain.

## Features
- Each instance of blockchain server will have its own wallet of initial balance 500, along with a private and public key.
    - User could then create transactions specifying the recipient address and the amount.
    - Transactions created would be signed with the wallet's private key and submitted to the blockchain's transaction pool.
- A new block can be added to the blockchain when a blockchain server successfully "mines" a new block.
    - For each transaction in the transaction pool, the blockchain server that is doing the mining work will verify the transaction and for the valid transactions, it will be added into a new block in the blockchain.
- Instances of the application can be synchronized (blockchain state & transactions in transaction pool) with each other using Websockets. Refer to section below on how to achieve this.

## How to Run Program
Head to the `/cmd/blockchain` directory and enter the command `go run .`. 

There are also two environment variables, `HTTP_PORT` and `PEERS` that can be specified along with running `go run .`. 
- `HTTP_PORT` specifies the port that the current blockchain listens to
- `PEERS` specifies the addresses of the servers that the current blockchain has to establish websocket connections to.

Example, if we have one blockchain server A running on address `http://localhost:8080`, we could start another blockchain server B which could connect to server A to synchronize blockchain updates. The command to start server B could be something like `HTTP_PORT=8081 PEERS=localhost:8080 go run .`.

Swagger docs would also be available at `{hostname}/swagger/index.html`.

## How to Run All Tests
From the root directory, enter command `go clean -testcache && go test ./...` to run all tests in the project recursively.

## How to Update Swagger Docs
From the root directory of the project, run `swag init -d ./cmd/blockchain -o ./cmd/docs --parseDependency`.
