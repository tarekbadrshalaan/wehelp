-- Table: public."helprequest"

-- DROP TABLE public."helprequest";

CREATE TABLE public.helprequest
(
    "id" SERIAL PRIMARY KEY,
    "details" TEXT NOT NULL,
    "is_done" BOOLEAN,
    "created_at" BIGINT
);