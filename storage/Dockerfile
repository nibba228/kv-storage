FROM golang

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

CMD ["go", "run", "/app/main.go", "/app/db.go"]

EXPOSE 8080
