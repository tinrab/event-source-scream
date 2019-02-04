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

CREATE TABLE "screams"
(
  "id"         bigint,
  "body"       text,
  "created_at" timestamp,
  "user_id"    bigint,

  PRIMARY KEY ("id")
);
