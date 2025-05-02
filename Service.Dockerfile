# syntax=docker/dockerfile:1

FROM golang:1.24.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY Makefile .
COPY VERSION .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN CGO_ENABLED=0 GOOS=linux make

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/service /service

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/service"]
