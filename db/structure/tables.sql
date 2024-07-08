/*
 * roles
 */
DROP TABLE IF EXISTS roles;
CREATE TABLE roles (
  role_id BIGINT PRIMARY KEY,
  role VARCHAR(64) UNIQUE NOT NULL,
  tps NUMERIC NOT NULL,
  comments TEXT NOT NULL
);

/*
 * users
 */
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  user_id BIGSERIAL PRIMARY KEY,
  role_id BIGINT DEFAULT NULL,
  tps NUMERIC NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  access_start TIMESTAMPTZ NOT NULL,
  access_end TIMESTAMPTZ NOT NULL,
  modified_ts TIMESTAMPTZ NOT NULL,
  modified_user_id BIGINT NOT NULL
);
CREATE INDEX ON users(password);
CREATE INDEX ON users(role_id);
CREATE INDEX ON users(access_start);
CREATE INDEX ON users(access_end);