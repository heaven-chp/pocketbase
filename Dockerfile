FROM golang:1.23.4 AS build
WORKDIR /work
ENV CGO_ENABLED=0
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/gomod-cache go mod download
COPY ./ ./
RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go build -o main ./main.go

FROM scratch
COPY --from=build /work/main /
ENTRYPOINT ["/main"]
