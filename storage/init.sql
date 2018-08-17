CREATE TABLE IF NOT EXISTS state (
  id               BIGSERIAL                NOT NULL PRIMARY KEY,
  context_id       TEXT                     NOT NULL,
  created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
  state            INT                      NOT NULL,
  data             BYTEA,
  event            BYTEA,
  processing_error TEXT
);

CREATE INDEX IF NOT EXISTS state_context_id
  ON state (context_id);