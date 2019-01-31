CREATE TABLE "events"
(
  "id"           bigint,
  "aggregate_id" bigint,
  "kind"         varchar(32),
  "version"      bigint,
  "fired_at"     timestamp,
  "data"         text,

  PRIMARY KEY ("id")
);

CREATE TABLE "users"
(
  "id"         bigint,
  "name"       varchar(32),

  PRIMARY KEY ("id")
);
