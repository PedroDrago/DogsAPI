CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    username text NOT NULL,
    email text NOT NULL,
    birth_year int NOT NULL CHECK(birth_year > 1980 AND birth_year <= EXTRACT(YEAR FROM NOW())),
    address text NOT NULL,
    phone text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    admin bool DEFAULT false
);
