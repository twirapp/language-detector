FROM golang:1.24-alpine as builder
WORKDIR /app

RUN apk add --no-cache --update \
    gcc \
    musl-dev \
    g++ \
    wget

COPY go.mod go.sum ./
RUN go mod download

RUN wget https://dl.fbaipublicfiles.com/fasttext/supervised-models/lid.176.bin

COPY . .

RUN CGO_ENABLED=1 go build -o /app/twir_application ./cmd/main.go

FROM alpine:3.21

RUN apk add --no-cache \
    libstdc++ \
    libgcc

COPY --from=builder /app/twir_application /bin/twir_application
COPY --from=builder /app/lid.176.bin /app/lid.176.bin
CMD ["/bin/twir_application", "-modelpath", "/app/lid.176.bin"]
