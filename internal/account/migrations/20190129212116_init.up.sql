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

CREATE TABLE "accounts"
(
  "id"         bigint,
  "title"      varchar(32),
  "created_at" timestamp,

  PRIMARY KEY ("id")
);
