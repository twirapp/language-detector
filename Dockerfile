FROM golang:1.24-alpine as builder
WORKDIR /app

RUN apk add upx

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/twir_application ./cmd/main.go && \
    upx -9 -k /app/twir_application

FROM alpine:latest
RUN apk add curl
COPY --from=builder /app/twir_application /bin/twir_application
CMD ["/bin/twir_application"]
