BEGIN;


CREATE TABLE IF NOT EXISTS word
(
    id          UUID PRIMARY KEY,
    word        VARCHAR(80) UNIQUE                    NOT NULL,
    translation VARCHAR(80)                           NOT NULL,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);


COMMIT;