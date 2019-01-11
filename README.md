# Pop

Officially developing Popcorn 2.0.

## Architecture

### Requirements

- Golang 1.11+
- Node 10+
- Docker

## Local Development Guide

### gRPC

Node.js API server communicates with recommender server through gRPC call. Whenever you modify the
proto files, you must run the shell script to recompile your proto outputs.

    ./protoc.sh

The shell script runs a Docker command to pull a `protoc` docker image and execute the image. `protoc`
is a protobuffer compiler that converts `.pb` files into the respective language files e.g. Golang
and JavaScript.

### API

You need to run `npm install` before you run `node index.js`

### Recommender

Simply run `go install && recommender`, we don't need `dep ensure` anymore.

## Deployment Guide

### Docker Compose

Each project has its `Dockerfile` which specifies the steps to build an image. We use `docker-compose`
to bring up all the images at once. If you don't have `docker-compose`, please download and install
it from [Docker.com][1].

Navigate to `popcorn/` directory and run

    docker-compose up --build

It will build all the images and bring up the containers. It should look like the following.

```text
Creating network "popcorn_default" with the default driver
Building backend
Step 1/8 : FROM golang:1.11.4
 ---> 343df9d12b7b
Step 2/8 : LABEL authors="Calvin Feng"
 ---> Running in 1fdb31313585
Removing intermediate container 1fdb31313585
 ---> af2401a1a5b5
Step 3/8 : COPY . /go/src/recommender
 ---> 009ef3fa9d72
Step 4/8 : WORKDIR /go/src/recommender
 ---> Running in 8f1950c9fb46
Removing intermediate container 8f1950c9fb46
 ---> 8f0437825ef4
Step 5/8 : ENV GO111MODULE=on
 ---> Running in 19be47aedef6
Removing intermediate container 19be47aedef6
 ---> 2f4adbe50a8a
Step 6/8 : EXPOSE 8080
 ---> Running in b24722826014
Removing intermediate container b24722826014
 ---> 8b20f818a61f
Step 7/8 : RUN go install
 ---> Running in 70a9d0770370
```

Once all the images are built, it will launch the containers.

```text
Creating recommender ... done
Creating api         ... done
Attaching to recommender, api
recommender | INFO[2019-01-11T20:36:02Z] Golang gRPC server is listening and serving on port 8080
api         | Node API server is serving and listening on port 8000
```

If you haven't changed any source code, there is no need to rebuild the image every time you run
`docker-compose`. You can simply do the following which is much faster to bring up the containers.

    docker-compose up

Now if you want to kill your containers, simply run the following in a seperate terminal.

    docker-compose down

```text
Stopping api         ... done
Stopping recommender ... done
Removing api         ... done
Removing recommender ... done
Removing network popcorn_default
```

[1]:https://docs.docker.com/compose/install/