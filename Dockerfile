FROM golang:1.13
RUN mkdir /app
WORKDIR /app
ADD main.go main.go
RUN go build -o app main.go
ENTRYPOINT ["/app/app"]
