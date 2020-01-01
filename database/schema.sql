CREATE TABLE schema_migrations (
    version VARCHAR(14) UNIQUE NOT NULL
);

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    provider_id TEXT NOT NULL,
    provider    TEXT,
    name        TEXT UNIQUE,
    real_name   TEXT,
    email       TEXT
);

CREATE TABLE habits (
    id      SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE activities (
    id       SERIAL PRIMARY KEY,
    habit_id INTEGER REFERENCES habits(id),
    created  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

