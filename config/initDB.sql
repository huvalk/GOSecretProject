-- create database manualy before this script

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL DEFAULT '',
    password TEXT NOT NULL DEFAULT '',
    phone TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS session (
    user_id INT NOT NULL,
    session_id TEXT NOT NULL DEFAULT '' UNIQUE
)