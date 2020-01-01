CREATE TABLE schema_migrations (
	version	VARCHAR(14) UNIQUE NOT NULL
);

CREATE TABLE account (
	user_id		SERIAL PRIMARY KEY,
	username	VARCHAR(50) UNIQUE NOT NULL,
	password	VARCHAR(50) NOT NULL,
	email		VARCHAR(355) UNIQUE NOT NULL,
	created_on	TIMESTAMP NOT NULL,
	last_login	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX ON account(last_login DESC);
CREATE INDEX ON account(last_login, last_login DESC);

CREATE TABLE role (
	role_id		SERIAL PRIMARY KEY,
	role_name	VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE account_role (
	user_id		INTEGER REFERENCES account(user_id),
	role_id		INTEGER REFERENCES role(role_id),
	grant_date	TIMESTAMP,
	PRIMARY KEY (user_id, role_id)
);

CREATE TABLE all_types (
	a_serial	SERIAL,
	a_boolean	BOOLEAN,
	a_char		CHAR(10),
	a_varchar	VARCHAR(10),
	a_int		INTEGER,
	a_ts		TIMESTAMP,
	a_ts_tz		TIMESTAMP WITH TIME ZONE,
	a_text		TEXT
);

CREATE TABLE users (
	provider	TEXT,
	name		TEXT,
	real_name	TEXT DEFAULT 'pickle',
	email		TEXT,
	PRIMARY KEY (provider, name)
);

