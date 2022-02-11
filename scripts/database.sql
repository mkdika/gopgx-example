CREATE DATABASE gopgx_example;

DROP TABLE IF EXISTS public.members;
CREATE TABLE IF NOT EXISTS public.members (
    id              BIGINT PRIMARY KEY,
    email           VARCHAR(100) NOT NULL,
    display_name    VARCHAR(100),
    description     TEXT,
    lvl             INTEGER,
    join_date       DATE,
    loyalty_point   NUMERIC DEFAULT 0.0,
    active          BOOLEAN DEFAULT false,
    created_at      TIMESTAMP(6) WITHOUT TIME ZONE NOT NULL
);

DROP SEQUENCE IF EXISTS public.members_id_seq;
CREATE SEQUENCE IF NOT EXISTS public.members_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.members_id_seq OWNED BY public.members.id;

