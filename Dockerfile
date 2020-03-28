FROM golang:1.13-alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
#COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags='-s -w -extldflags "-static"' -o ./battleship ./main.go

# final stage
FROM golang:1.13-alpine
WORKDIR /app

COPY --from=builder /app/battleship .

CMD ["./battleship -client=false"]
