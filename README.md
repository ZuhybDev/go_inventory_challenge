# go_inventory_challenge

A small, self-contained Go challenge implementing a simple inventory service. This repository is intended as a lightweight exercise to demonstrate basic Go project structure, build/run workflow, and simple CRUD operations for inventory items.

## Features

- Basic inventory item model (ID, name, quantity, price)
- Simple in-memory storage (easy to replace with a DB)
- HTTP API for CRUD operations (common endpoints listed below)
- Tests (if present) can be run with `go test` across the module

## Prerequisites

- Go 1.20+ installed (older versions may work but are untested)
- Git (to clone the repo)

## Get started

1. Clone the repository:

   git clone https://github.com/ZuhybDev/go_inventory_challenge.git
   cd go_inventory_challenge

2. Build (to produce a binary):

   go build -o bin/inventory ./...

3. Run (development mode):

   # Run main package (if main is at repository root)
   go run .

   # Or run the built binary
   ./bin/inventory

Note: If the project provides a `cmd/` directory with a specific entrypoint, replace `.` with the appropriate path (for example `./cmd/server`).

## Configuration

The service may read configuration from environment variables. Typical options include:

- PORT — HTTP port to listen on (default: 8080)
- DATA_FILE — optional file path to persist data between runs

Adjust according to the implementation in the repository.

## API (example)

If the challenge exposes an HTTP API, typical endpoints are:

- GET  /items         — list all items
- POST /items         — create a new item (JSON body)
- GET  /items/{id}    — fetch one item by ID
- PUT  /items/{id}    — update an existing item
- DELETE /items/{id}  — remove an item

Example curl requests:

Create an item:

  curl -X POST -H "Content-Type: application/json" \
    -d '{"name":"Widget","quantity":10,"price":9.99}' http://localhost:8080/items

List items:

  curl http://localhost:8080/items

## Tests

Run unit tests with:

  go test ./...

If there are test files included in the repo, this will run them and report results.

## Contributing

This repository is a small challenge project — contributions are welcome. If you want to improve the code, add tests, or provide persistence/backing stores, please open a pull request.

To report bugs or request features, open an issue:

  https://github.com/ZuhybDev/go_inventory_challenge/issues

## License

Add a license file (e.g. MIT) if you want to make the project open source. If no license is present, the repository defaults to "All rights reserved" for the owner.

## Author

ZuhybDev — created as a small Go inventory challenge.
