CREATE TABLE event (
    uuid        UUID PRIMARY KEY,
    log_uuid    UUID NOT NULL,
    event_type  VARCHAR,
    event_txt  VARCHAR,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    deleted_at  timestamp default null,
    CONSTRAINT FK_log_uuid FOREIGN KEY (log_uuid) REFERENCES log(uuid)
)