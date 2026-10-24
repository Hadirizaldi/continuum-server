CREATE TABLE upload_sessions (
id            VARCHAR(36) PRIMARY KEY,
filename      VARCHAR(255) NOT NULL,
total_size    BIGINT       NOT NULL,
mime_type     VARCHAR(127) NOT NULL,
user_id       VARCHAR(255),
chunk_size    BIGINT       NOT NULL,
total_chunks  INT          NOT NULL,
status        VARCHAR(20)  NOT NULL DEFAULT 'pending',
created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),
last_activity_at TIMESTAMPTZ NOT NULL DEFAULT now(),
expires_at      TIMESTAMPTZ NOT NULL,
completed_at    TIMESTAMPTZ,
final_path      TEXT,
file_url        TEXT
);


CREATE INDEX idx_sessions_status ON upload_sessions (status);
CREATE INDEX idx_sessions_expires_at ON upload_sessions (expires_at);
CREATE INDEX idx_sessions_user_id ON upload_sessions (user_id);


CREATE TABLE upload_chunks (
    session_id   VARCHAR(36) NOT NULL REFERENCES upload_sessions(id) ON DELETE CASCADE,
    chunk_index  INT         NOT NULL,
    chunk_size   BIGINT      NOT NULL,
    offset_start BIGINT      NOT NULL,
    offset_end   BIGINT      NOT NULL,
    storage_path TEXT        NOT NULL,
    PRIMARY KEY (session_id, chunk_index)
);
