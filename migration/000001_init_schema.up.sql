CREATE TABLE IF NOT EXISTS  comments (
    id serial PRIMARY KEY ,
    slug varchar,
    body varchar,
    author varchar
)