CREATE TABLE users
(
    id            UUID PRIMARY KEY,
    login         TEXT,
    token         TEXT NOT NULL,
    uid           BIGINT,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rooms
(
    id         UUID PRIMARY KEY,
    name       TEXT NOT NULL UNIQUE,
    password   TEXT NOT NULL,
    owner_id   UUID REFERENCES users (id) ON DELETE CASCADE,
    users_count BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tracks
(
    id         UUID PRIMARY KEY,
    room_id    UUID REFERENCES rooms (id) ON DELETE CASCADE,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE,
    track_data JSONB NOT NULL,
    status     TEXT      DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);