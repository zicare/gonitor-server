/*
 * apps
 */
CREATE TYPE APP_TYPE AS ENUM ('WEB', 'ANDROID', 'IOS');
CREATE TABLE apps (
  app_id BIGSERIAL PRIMARY KEY,
  type APP_TYPE NOT NULL,
  key VARCHAR(36) UNIQUE NOT NULL,
  origin VARCHAR(128) UNIQUE NOT NULL,
  app VARCHAR(128) UNIQUE NOT NULL,
  description TEXT NOT NULL
);

/*
 * roles
 */
CREATE TABLE roles (
  role_id BIGINT PRIMARY KEY,
  role VARCHAR(64) UNIQUE NOT NULL,
  tps NUMERIC NOT NULL,
  comments TEXT NOT NULL
);

/*
 * agents
 */
CREATE TABLE agents (
  agent_id BIGSERIAL PRIMARY KEY,
  role_id BIGINT DEFAULT NULL REFERENCES roles(role_id) ON DELETE RESTRICT,
  tps NUMERIC NOT NULL,
  agent VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  access_start TIMESTAMPTZ NOT NULL,
  access_end TIMESTAMPTZ NOT NULL,
  params JSON NOT NULL
);
CREATE INDEX ON agents(password);
CREATE INDEX ON agents(role_id);
CREATE INDEX ON agents(access_start);
CREATE INDEX ON agents(access_end);

/*
 * users
 */
CREATE TABLE users (
  user_id BIGSERIAL PRIMARY KEY,
  role_id BIGINT DEFAULT NULL REFERENCES roles(role_id) ON DELETE RESTRICT,
  tps NUMERIC NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  access_start TIMESTAMPTZ NOT NULL,
  access_end TIMESTAMPTZ NOT NULL
);
CREATE INDEX ON users(password);
CREATE INDEX ON users(role_id);
CREATE INDEX ON users(access_start);
CREATE INDEX ON users(access_end);

/*
 * pins
 */
CREATE TABLE pins (
  pin_id BIGSERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL REFERENCES users(email) ON DELETE CASCADE,
  pin VARCHAR(5) NOT NULL,
  created TIMESTAMPTZ NOT NULL,
  expiration TIMESTAMPTZ NOT NULL,
  UNIQUE (email, created)
);
CREATE INDEX ON pins(email, pin);

/*
 * endpoints
 */
CREATE TYPE METHOD AS ENUM ('GET', 'HEAD', 'POST', 'PUT', 'PATCH', 'DELETE');
CREATE TABLE endpoints (
  endpoint_id BIGINT PRIMARY KEY,
  method METHOD NOT NULL,
  route VARCHAR(255) NOT NULL
);

/*
 * grants
 */
CREATE TABLE grants (
  grant_id BIGSERIAL PRIMARY KEY,
  endpoint_id BIGINT NOT NULL REFERENCES endpoints(endpoint_id) ON DELETE CASCADE,
  role_id BIGINT NOT NULL REFERENCES roles(role_id) ON DELETE CASCADE,
  start_date TIMESTAMPTZ NOT NULL,
  end_date TIMESTAMPTZ NOT NULL,
  UNIQUE(endpoint_id,role_id)
);
CREATE INDEX ON grants(endpoint_id);
CREATE INDEX ON grants(role_id);
CREATE INDEX ON grants(start_date);
CREATE INDEX ON grants(end_date);

/*
 * host_groups
 */
CREATE TABLE host_groups (
  host_group_id BIGSERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  description TEXT NOT NULL,
  params JSON NOT NULL
);

/*
 * hosts
 */
CREATE TABLE hosts (
  host_id BIGSERIAL PRIMARY KEY,
  hostname VARCHAR(255) NOT NULL,
  ip_address INET NOT NULL,
  description TEXT NOT NULL,
  params JSON NOT NULL,
  UNIQUE (hostname, ip_address)
);
CREATE INDEX ON hosts(ip_address);

/*
 * lnk_host_groups_hosts
 */
CREATE TABLE lnk_host_groups_hosts (
  host_group_id BIGINT NOT NULL REFERENCES host_groups(host_group_id) ON DELETE CASCADE,
  host_id BIGINT NOT NULL REFERENCES hosts(host_id) ON DELETE CASCADE,
  PRIMARY KEY (host_group_id,host_id)
);
CREATE INDEX ON lnk_host_groups_hosts(host_id);

/*
 * items
 */
CREATE TYPE DATATYPE AS ENUM ('DECIMAL', 'INTEGER', 'HEX', 'BINARY', 'TIMESTAMP');
CREATE TABLE items (
  item_id BIGSERIAL PRIMARY KEY,
  item VARCHAR(255) UNIQUE NOT NULL,
  unit VARCHAR(255) NOT NULL,
  data_type DATATYPE NOT NULL,
  description TEXT NOT NULL,
  params JSON NOT NULL
);

/*
 * lnk_hosts_items
 */
CREATE TABLE lnk_hosts_items (
  host_id BIGINT NOT NULL REFERENCES hosts(host_id) ON DELETE CASCADE,
  item_id BIGINT NOT NULL REFERENCES items(item_id) ON DELETE CASCADE,
  schedule VARCHAR(64) NOT NULL,
  params JSON NOT NULL,
  PRIMARY KEY (host_id,item_id)
);
CREATE INDEX ON lnk_hosts_items(item_id);

/*
 * events
 */
CREATE TABLE events (
  event_id BIGSERIAL PRIMARY KEY,
  host_id BIGINT NOT NULL REFERENCES hosts(host_id) ON DELETE RESTRICT,
  item_id BIGINT NOT NULL REFERENCES items(item_id) ON DELETE RESTRICT,
  data NUMERIC NOT NULL,
  datetime TIMESTAMPTZ NOT NULL
);
CREATE INDEX ON events USING BRIN(host_id, item_id, datetime);