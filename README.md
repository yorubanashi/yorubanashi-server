- `main.go`: Entrypoint for the main server
- `cmd/gen/main.go`: Generate handlers for server based on function signatures
- `cmd/songconv/main.go`: Converts song lyrics in text format into YAML
  - Currently relies on having [OpenCC](https://github.com/BYVoid/OpenCC) cloned locally
  - Hint: run with `go run cmd/songconv/*.go`
