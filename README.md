# Simple-Blockchain
A blockchain project built to gain understanding on the fundamental workings of blockchain.

## Features
- Instances of the application can be synchronized (blockchain state & transactions in transaction pool) with each other using Websockets. Refer to section below on how to achieve this.
- Each spawned server has its own wallet, with a private/public key. Transactions can then be created and submitted to the transaction pool.
- A new block can be added to the blockchain when valid transactions from the transaction pool have been pulled and successfully added into a new block (a.k.a the mining process).


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
