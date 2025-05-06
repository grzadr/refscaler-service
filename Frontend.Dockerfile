# Frontend.Dockerfile
FROM golang:1.24.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum Makefile VERSION ./

COPY cmd/ ./cmd/
COPY internal/ ./internal/
COPY assets/ ./assets/
COPY assets/views/partials ./assets/views/partials

RUN CGO_ENABLED=0 GOOS=linux make build-frontend

FROM gcr.io/distroless/base-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/bin/frontend /frontend
COPY --from=build-stage /app/assets/ /assets/
COPY --from=build-stage /app/assets/views/partials /assets/views/partials
ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/frontend"]
