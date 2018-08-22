CREATE TABLE users(
    id          bigserial NOT NULL,
    login       text UNIQUE NOT NULL,
    pass        bytea NOT NULL,
    created_at  timestamp DEFAULT(now() at time zone 'utc') NOT NULL,
    updated_at  timestamp DEFAULT(now() at time zone 'utc') NOT NULL,
    deleted_at  timestamp,
    status      text NOT NULL default 'active', -- it should be good idea to make this some enum type
    role        text NOT NULL default 'user', -- this too

    PRIMARY KEY(id)
);
