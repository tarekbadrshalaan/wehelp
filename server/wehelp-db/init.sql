
CREATE TABLE public.address
(
    "id" SERIAL PRIMARY KEY,
    "address" TEXT,
    "latlng" TEXT,
    "city" TEXT,
    "is_active" BOOLEAN,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL
);


CREATE TABLE public.user
(
    "id" SERIAL PRIMARY KEY,
    "email" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "password_md5" VARCHAR (32)  NOT NULL,
    "phone" TEXT,
    "bio" TEXT,
    "address" INT,
    "is_active" BOOLEAN,
    "created_at" BIGINT NOT NULL,
    "updated_at" BIGINT NOT NULL,
    FOREIGN KEY (address) REFERENCES public.address (id)
);
