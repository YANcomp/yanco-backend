-- +goose Up
BEGIN;

SET
    statement_timeout = 0;
SET
    client_encoding = 'UTF8';
SET
    standard_conforming_strings = ON;
SET
    check_function_bodies = FALSE;
SET
    client_min_messages = WARNING;
SET
    search_path = public, extensions;
SET
    default_tablespace = '';
SET
    default_with_oids = FALSE;

-- EXTENSIONS --
CREATE
    EXTENSION IF NOT EXISTS pgcrypto;
CREATE
    EXTENSION IF NOT EXISTS citext;

-- DROP TABLE --
DROP TABLE IF EXISTS banners CASCADE;

create table banners
(
    id          serial PRIMARY KEY,
    uuid        UUID                     DEFAULT gen_random_uuid(),
    active      boolean                  DEFAULT false,
    sort        integer                  DEFAULT 500,
    header      VARCHAR(250) NOT NULL CHECK ( header <> '' ),
    description text                     DEFAULT '',
    targettype  text                     DEFAULT '',
    target      text                     DEFAULT '',
    image       text                     DEFAULT '',
    imagemobile text                     DEFAULT '',
    iscatalog   boolean                  DEFAULT false,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

COMMIT;

-- +goose Down
drop table banners;