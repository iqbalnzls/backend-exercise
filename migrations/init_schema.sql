CREATE DATABASE backend_exercise;

CREATE USER backend_exercise_user WITH PASSWORD 'backend_exercise_password';

CREATE SCHEMA backend_exercise;

CREATE TABLE backend_exercise.listings
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER     NOT NULL,
    price        INTEGER     NOT NULL,
    listing_type VARCHAR(64) NOT NULL,
    created_at   BIGINT      NOT NULL DEFAULT (extract(epoch FROM now()) * 1000000)::BIGINT,
    updated_at   BIGINT      NOT NULL DEFAULT (extract(epoch FROM now()) * 1000000)::BIGINT,

    CONSTRAINT user_id_fkey FOREIGN KEY (user_id) REFERENCES backend_exercise.users (id) ON DELETE CASCADE

);

CREATE TABLE backend_exercise.users
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(256) NOT NULL,
    created_at BIGINT       NOT NULL DEFAULT (extract(epoch FROM now()) * 1000000)::BIGINT,
    updated_at BIGINT       NOT NULL DEFAULT (extract(epoch FROM now()) * 1000000)::BIGINT
)