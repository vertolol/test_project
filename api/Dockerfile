FROM golang:latest

RUN mkdir /api
WORKDIR /api

COPY /api /api
RUN apt -y update && apt -y install git


RUN go build

CMD ["./api"]
