# Recommender 

## Docker

    docker build -t recommender .
    docker run -p 8080:8080 recommender

If you use `-P` to publish all exposed ports, they will get mapped to random port numbers. I
recommend using `-p <exposed>:<publish_to>` to specify the desired port to publish to.