CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(55) UNIQUE NOT NULL,
    email VARCHAR(55) UNIQUE NOT NULL
);