FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o apprise-api ./cmd/apprise-mvp

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/apprise-api /app/apprise-api
COPY --from=builder /app/docs /app/docs

EXPOSE 8080

CMD ["/app/apprise-api"]