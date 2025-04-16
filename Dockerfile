# syntax=docker/dockerfile:1

FROM golang:1.24.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/

RUN go test -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /service cmd/service/main.go

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /service /service

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/service"]
