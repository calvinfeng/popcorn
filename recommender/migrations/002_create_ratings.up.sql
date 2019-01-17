CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    movie_id INTEGER REFERENCES movies(id),
    email VARCHAR(256),
    value REAL 
);