CREATE TABLE users (
    id serial primary key,
    fullname text,
    username text unique not null,

    token text,
    token_expire_date timestamp,

    last_login timestamp,

    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);