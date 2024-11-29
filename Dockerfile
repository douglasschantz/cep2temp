FROM golang:1.21 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cep2temp cmd/apid/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/cep2temp .
COPY ./configs/config.env ./configs/config.env
ENTRYPOINT ["./cep2temp"]
