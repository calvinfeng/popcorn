# Pop

Officially developing Popcorn 2.0.

## Architecture

## Developer Guide

### gRPC

Node.js API server communicates with Golang microservice through gRPC call. Whenever you modify the
protobuf files, you must run the shell script to recompile your proto outputs. The shell script runs
a Docker command to pull a `protoc` docker image and execute the image.

    ./protoc.sh
