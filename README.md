# Simple-Blockchain

## How to Run Program
Head to the `/cmd/blockchain` directory and enter the command `go run .`. 

There are also two environment variables, `HTTP_PORT` and `PEERS` that can be specified along with running `go run .`. 
- `HTTP_PORT` specifies the port that the current blockchain listens to
- `PEERS` specifies the addresses of the servers that the current blockchain has to establish websocket connections to.

Example, if we have one blockchain server A running on address `http://localhost:8080`, we could start another blockchain server B which could connect to server A to synchronize blockchain updates. The command to start server B could be something like `HTTP_PORT=8081 PEERS=localhost:8080 go run .`.
## How to Run All Tests
From the root directory, enter command `go clean -testcache && go test ./...` to run all tests in the project recursively.