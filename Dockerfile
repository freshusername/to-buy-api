FROM golang:1.20.1 as builder

ARG DB_HOST
ARG DB_NAME
ARG DB_PASSWORD
ARG DB_PORT
ARG DB_USER
ARG UI_HOST

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
#COPY .env.local /app - works only locally

ENV DB_HOST=${DB_HOST}
ENV DB_NAME=${DB_NAME}
ENV DB_PASSWORD=${DB_PASSWORD}
ENV DB_PORT=${DB_PORT}
ENV DB_USER=${DB_USER}
ENV UI_HOST=${UI_HOST}

EXPOSE 3000
CMD ["/app/server"]