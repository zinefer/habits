CREATE TABLE users (
    provider TEXT,
    name TEXT,
    real_name TEXT default 'kosher',
	email TEXT,
    PRIMARY KEY (provider, name)
);