FROM golang:1.11.4 
LABEL authors="Calvin Feng"

COPY . /go/src/recommender
WORKDIR /go/src/recommender

ENV GO111MODULE=on

ENV GCP_PORT=8080
ENV DEFAULT_PORT=8081

EXPOSE 8080
EXPOSE 8081

RUN go install
RUN recommender migrate

CMD recommender serve