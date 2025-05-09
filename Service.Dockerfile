# syntax=docker/dockerfile:1

FROM golang:1.24.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum Makefile VERSION ./

COPY cmd/ ./cmd/
COPY internal/ ./internal/

RUN CGO_ENABLED=0 GOOS=linux make build-service

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/service /service
ENV PORT=3000
EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/service"]
