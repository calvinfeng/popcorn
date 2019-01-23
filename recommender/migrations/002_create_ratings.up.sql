CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    movie_id INTEGER REFERENCES movies(id),
    user_email VARCHAR(256),
    value REAL,
    UNIQUE (movie_id, user_email)
);

CREATE INDEX ON ratings(user_email);

CREATE TABLE preferences (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    user_email VARCHAR(256),
    value DOUBLE PRECISION[]   
);

CREATE UNIQUE INDEX ON preferences(user_email);