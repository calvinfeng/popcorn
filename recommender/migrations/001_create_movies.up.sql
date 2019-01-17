CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    imdb_id VARCHAR(256),
    tmdb_id VARCHAR(256),
    title VARCHAR(512),
    year SMALLINT,
    num_rating INTEGER,
    imdb_rating REAL,
    average_rating REAL,
    feature DOUBLE PRECISION[],
    cluster SMALLINT,
    nearest_clusters SMALLINT[],
    farthest_clusters SMALLINT[]
);

CREATE INDEX ON movies(title);
CREATE INDEX ON movies(year);
CREATE INDEX ON movies(cluster);


CREATE TABLE movie_details (
    imdb_id VARCHAR(512) PRIMARY KEY,
    detail BYTEA
);

CREATE TABLE movie_trailers (
    imdb_id VARCHAR(512) PRIMARY KEY,
    trailer BYTEA
);