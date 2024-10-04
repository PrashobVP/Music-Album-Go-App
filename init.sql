CREATE DATABASE IF NOT EXISTS music;

USE music;

CREATE TABLE IF NOT EXISTS albums (
    id varchar(255) primary key,
    title varchar(255) not null,
    artisit varchar(255) not null,
    price decimal(10, 2) not null
);
