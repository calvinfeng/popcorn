# Recommender

## Docker

    docker build -t recommender . --network=host
    docker run -p 8081:8081 recommender --network=host

If you use `-P` to publish all exposed ports, they will get mapped to random port numbers. I
recommend using `-p <exposed>:<publish_to>` to specify the desired port to publish to.

## Database

For development, enter Postgres shell

    psql postgres

Ccreate a role if you don't have it already

```sql
CREATE ROLE popcorn superuser login;
ALTER USER popcorn WITH PASSWORD 'popcorn';
```

Create a database

```sql
CREATE DATABASE popcorn with OWNER=popcorn;
```

## Available Commands

### Run server

    recommender serve

### Run migration

    recommender migrate

### Drop database

    recommender dropdb

### Seed

    recommender seed

## Datasets

In order for seeding to work, download the data from movies lens using the following command.

    cd popcorn/recommender/datasets
    wget http://files.grouplens.org/datasets/movielens/ml-latest-small.zip

And then unzip it into your `datasets/` directory.