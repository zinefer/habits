CREATE TABLE schema_migrations (
	version	VARCHAR(14) UNIQUE NOT NULL
);

CREATE TABLE users (
	id			SERIAL PRIMARY KEY,
	provider_id TEXT NOT NULL,
	provider	TEXT,
	name		TEXT UNIQUE,
	real_name	TEXT,
	email		TEXT
);