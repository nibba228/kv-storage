FROM grpc/go:latest

WORKDIR /storage

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

CMD go run main.go

EXPOSE 8080
