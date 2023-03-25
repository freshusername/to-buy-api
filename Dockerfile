FROM golang:1.20.1 as builder

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

#retrieve application dependencies, copy go.mod & go.sum
COPY go.mod ./
COPY go.sum ./
RUN go mod download

#copy local code to the container image
COPY . ./ 

#print packages when building the binary
RUN CGO_ENABLED=0 go build -v -o server

FROM alpine:latest as deploy
COPY --from=builder /app/server /app/server
COPY .env /app
EXPOSE 3000
CMD ["/app/server"]