# TSS Implementation

This project implements a Threshold Signature Scheme (TSS) using gRPC and Redis.

## Project Structure

```
tss-impl
├── client
│ └── client.go
├── proto
│ └── [proto files]
├── utils
│ └── utils.go
├── go.mod
├── go.sum
└── README.md
```

## Prerequisites

- Go 1.23.4 or later
- PostgreSQL
- Redis

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/tss-impl.git
   cd tss-impl
   ```

2. Copy the example environment file and update it with your configuration:

   ```sh
   cp .env.example .env
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Run database migrations:
   ```sh
   make migrate-up
   ```

## Running the Project

1. Start the server:

   ```sh
   make run
   ```

2. Run the client:
   ```sh
   go run client/client.go
   ```

## Project Components

### Client

The client initiates key generation and signing processes. It communicates with the server using gRPC.

#### Key Generation

The client sends a `NotifyAction` request to the server to initiate the key generation process. The server responds with the generated key shares, which are then saved to a file.

#### Signing

The client loads the key shares from the file and sends a `NotifyAction` request to the server to initiate the signing process. The server responds with the signature.

### Server

The server handles key generation and signing requests. It uses Redis for inter-node communication and PostgreSQL for storing encrypted shares.

#### Key Generation

The server receives a `NotifyAction` request from the client to initiate the key generation process. It generates the key shares and publishes them to a Redis channel.

#### Signing

The server receives a `NotifyAction` request from the client to initiate the signing process. It retrieves the key shares from the database and performs the signing operation.

### Configuration

Configuration is managed using environment variables. See [.env.example](http://_vscodecontentref_/2) for the required variables.

### Database

The database schema is managed using SQL migrations. See the [migrations](http://_vscodecontentref_/3) directory for the migration files.

### Utilities

Utility functions for file operations, encryption, and compression are located in the [utils](http://_vscodecontentref_/4) directory.

## License

This project is licensed under the MIT License.
