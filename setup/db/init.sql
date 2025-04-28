-- init.sql
CREATE TABLE IF NOT EXISTS public.users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name  VARCHAR(255)
);
