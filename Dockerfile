FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o git-committer ./cmd/main.go

CMD ["./git-committer"]