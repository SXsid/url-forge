FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.sum go.mod  ./
RUN go mod download
COPY . .
RUN go build -o server  ./cmd/server/main.go

FROM  scratch 
COPY --from=builder  /app/server /server
EXPOSE 8080
CMD [ "/server" ]

