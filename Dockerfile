FROM golang:latest as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
#RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 go build -o /tobuyapi
#RUN make build
RUN ls

#RUN chmod +x /app/tobuyapi

# run in small image
FROM alpine:latest
RUN apk add --no-cache ca-certificates

RUN mkdir /app

COPY --from=builder /tobuyapi /app

EXPOSE 3000
CMD ["/app/tobuyapi"]