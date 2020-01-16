CREATE TABLE schema_migrations (
    version VARCHAR(14) UNIQUE NOT NULL
);

CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    provider_id TEXT NOT NULL,
    provider    TEXT,
    name        TEXT UNIQUE,
    real_name   TEXT,
    email       TEXT,
    created     TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(provider_id, provider)
);

CREATE TABLE habits (
    id      SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    name    TEXT NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ON habits(user_id);

CREATE TABLE activities (
    id        SERIAL PRIMARY KEY,
    habit_id  INTEGER REFERENCES habits(id),
    moment    TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    time_zone INTEGER DEFAULT 0,
    created   TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ON activities(habit_id, moment);

