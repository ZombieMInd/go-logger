CREATE TABLE log (
    uuid        UUID PRIMARY KEY,
    ip          VARCHAR(64),
    timestamp timestamp,
    user_uuid   UUID,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    deleted_at  timestamp default null
)