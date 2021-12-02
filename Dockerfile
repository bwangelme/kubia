# docker build -t bwangel/qae_kubia_$(date +%Y_%m_%d_%H%M) --build-arg app=kubia .

FROM golang:1.17 as builder

ARG app

RUN mkdir /code
WORKDIR /code
COPY . /code
RUN go build -o ${app} .

FROM golang:1.17

ARG app

COPY --from=builder /code/${app} /go/bin
