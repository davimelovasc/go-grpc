# go-grpc Study Project

This is a simple study project created to gain a better understanding of how gRPC works with Go. The project aims to demonstrate the basic concepts of gRPC, utilizing both a gRPC client and server, conveniently located within the same repository for ease of understanding.


## Project Structure

The project's repository is organized as follows:

- `cmd/server/`: Contains the implementation of the gRPC server.
- `cmd/client/`: Includes the implementation of the gRPC client.
- `proto/`: This directory holds the Protocol Buffers (protobuf) definition files used to define the gRPC service and messages.

## How to Use the Project

To use this project, follow these steps:

1. Clone the repository to your local machine using the following command:
   ```
   git clone https://github.com/davimelovasc/go-grpc.git
   ```

2. Navigate to the project directory:
   ```
   cd go-grpc
   ```

3. Build and run the gRPC server:
   ```
   go run server/main.go
   ```

4. Build and run the gRPC client (in a separate terminal window):
   ```
   go run client/main.go
   ```