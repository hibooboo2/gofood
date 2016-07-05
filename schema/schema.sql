DROP TABLE IF EXISTS ingredients CASCADE;
DROP TABLE IF EXISTS store_ingredients CASCADE;
DROP TABLE IF EXISTS recipes CASCADE;
DROP TABLE IF EXISTS tags CASCADE;

DROP TYPE IF EXISTS measuring_unit;

CREATE TYPE measuring_unit AS ENUM ('g','oz','cups','tbs','tsp','lbs','kg');

CREATE TABLE ingredients(
	id SERIAL PRIMARY KEY,
	name VARCHAR(128) NOT NULL UNIQUE,
	serving_size INT NOT NULL,
	serving_unit measuring_unit NOT NULL,
	calories_per_serving INT NOT NULL
);

CREATE INDEX ON ingredients (id);
CREATE INDEX ON ingredients (name);


CREATE TABLE store_ingredients(
	id SERIAL PRIMARY KEY,
	ingredient_id  INT REFERENCES ingredients (id), 
	store VARCHAR(120) NOT NULL,
	price MONEY NOT NULL,
	currency VARCHAR(10) DEFAULT 'USD' NOT NULL,
	package_size INT NOT NULL,
	package_size_unit measuring_unit NOT NULL
);

CREATE TABLE tags(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) UNIQUE
);

CREATE TABLE recipes(
	id SERIAL PRIMARY KEY,
	ingredient_ids INT[] NOT NULL,
	name VARCHAR(128) NOT NULL,
	steps TEXT NOT NULL,
	tags INT[]
);

