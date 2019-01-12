#!/bin/bash
docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l go -o recommender/pb
docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l node -o api/server/pb 