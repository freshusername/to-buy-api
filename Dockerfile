FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
COPY go.mod ./
#COPY go.sum ./
COPY . .

#RUN make build
#RUN go build -o /bin/app
#EXPOSE 3000

RUN go mod tidy
CMD ["/bin/app"]