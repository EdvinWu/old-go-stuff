CREATE SCHEMA htt;

CREATE TABLE htt.parent_accounts (

  id         BIGSERIAL NOT NULL,
  first_name VARCHAR   NOT NULL,
  last_name  VARCHAR   NOT NULL,
  email      VARCHAR   NOT NULL,
  login      VARCHAR   NOT NULL,
  password   VARCHAR   NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE htt.child_accounts (

  id         BIGSERIAL                                  NOT NULL,
  parent_id  BIGINT REFERENCES htt.parent_accounts (id) NOT NULL,
  first_name VARCHAR                                    NOT NULL,
  last_name  VARCHAR                                    NOT NULL,
  email      VARCHAR,
  login      VARCHAR                                    NOT NULL,
  password   VARCHAR                                    NOT NULL,
  points     VARCHAR DEFAULT 0,

  PRIMARY KEY (id)
);

CREATE TABLE htt.tasks (

  id          BIGSERIAL                                    NOT NULL,
  creator_id  BIGINT REFERENCES htt.parent_accounts (id)   NOT NULL,
  assignee_id BIGINT REFERENCES htt.child_accounts (id),
  name        VARCHAR                                      NOT NULL,
  description TEXT                                         NOT NULL,
  points      VARCHAR                                      NOT NULL,
  status      VARCHAR                                      NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE htt.goals (

  id          BIGSERIAL                                  NOT NULL,
  creator_id  BIGINT REFERENCES htt.parent_accounts (id) NOT NULL,
  name        VARCHAR                                    NOT NULL,
  description TEXT                                       NOT NULL,
  cost        VARCHAR                                    NOT NULL,
  status      VARCHAR                                    NOT NULL,

  PRIMARY KEY (id)
);
