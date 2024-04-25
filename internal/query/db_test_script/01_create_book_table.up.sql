-- +migrate Up
CREATE SCHEMA store;

CREATE TABLE store.books (
    id UUID NOT NULL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    author VARCHAR(50) NOT NULL,
    price DOUBLE PRECISION
);

