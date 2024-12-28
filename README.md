# pocketbase

## Overview

### Use
- build
  - `go build -o main main.go`
- run
  - `./main serve --config-file ./config/config.json`
- initial superuser
  - `admin@admin.com`/`admin123`

### Test and Coverage
- Test
  - `go clean -testcache && go test -cover ./...`
- Coverage
  - make coverage file
    - `go clean -testcache && go test -coverprofile=coverage.out -cover ./...`
  - convert coverage file to html file
    - `go tool cover -html=./coverage.out -o ./coverage.html`
