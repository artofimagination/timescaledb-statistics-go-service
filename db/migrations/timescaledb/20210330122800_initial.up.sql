-- +migrate Up
CREATE TABLE IF NOT EXISTS categories(
   id SERIAL PRIMARY KEY,
   name VARCHAR (50),
   description VARCHAR (300),
   created_at timestamp NOT NULL DEFAULT NOW(),
   updated_at timestamp NOT NULL DEFAULT NOW()
);

-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_stats_json(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data jsonb
);

-- +migrate Up
CREATE INDEX ON daily_stats_json (created_at DESC, category);

-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_stats_fp(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data double precision
);

-- +migrate Up
CREATE INDEX ON daily_stats_fp (created_at DESC, category);

-- +migrate Up
CREATE TABLE IF NOT EXISTS daily_stats_int(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data double precision
);

-- +migrate Up
CREATE INDEX ON daily_stats_int (created_at DESC, category);

-- +migrate Up
CREATE TABLE IF NOT EXISTS statistics_json(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data jsonb
);

-- +migrate Up
CREATE INDEX ON statistics_json (created_at DESC, category);

-- +migrate Up
CREATE TABLE IF NOT EXISTS statistics_fp(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data double precision
);

-- +migrate Up
CREATE INDEX ON statistics_fp (created_at DESC, category);

-- +migrate Up
CREATE TABLE IF NOT EXISTS statistics_int(
   created_at timestamp NOT NULL DEFAULT NOW() PRIMARY KEY,
   category smallint NOT NULL,
   FOREIGN KEY (category) REFERENCES categories(id),
   data bigint
);

-- +migrate Up
CREATE INDEX ON statistics_int (created_at DESC, category);