CREATE TABLE IF NOT EXISTS dogs (
    id BIGSERIAL PRIMARY KEY,
    name text NOT NULL,
    birth_year int NOT NULL CHECK(birth_year > 1980 AND birth_year <= EXTRACT(YEAR FROM NOW())),
    breed text NOT NULL,
    sex text NOT NULL,
    special_needs text[] NOT NULL,
    neutered bool NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    user_id bigint NOT NULL REFERENCES users on DELETE CASCADE,
    version integer NOT NULL DEFAULT 1
);
