# pocketbase

## Overview

### Use
- build
  - `go build -o main main.go`
- run
  - `./main serve --config-file ./config/config.json`
- initial superuser
  - `admin@admin.com`/`admin123`
- How to create a new collection
  1. Define a new schema by referencing the sample schema(`collections/schemas/sample.go`)
  2. Adding a new schema to the Get function(`collections/schemas/schemas.go`)

### Test and Coverage
- Test
  - `go clean -testcache && go test -cover ./...`
- Coverage
  - make coverage file
    - `go clean -testcache && go test -coverprofile=coverage.out -cover ./...`
  - convert coverage file to html file
    - `go tool cover -html=./coverage.out -o ./coverage.html`
