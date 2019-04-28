#!/bin/bash
docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l go -o recommender/protogo
docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l node -o api/server/protojs
docker run -v $(pwd):/defs namely/protoc-all:1.17_0 -d protos -l python -o tfrec/