FROM golang:1.11.4 
LABEL authors="Calvin Feng"

COPY . /go/src/recommender
WORKDIR /go/src/recommender

ENV GO111MODULE=on

EXPOSE 8080

RUN go install

CMD recommender