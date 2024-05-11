CREATE TABLE gallery (
    id VARCHAR(36) PRIMARY KEY,
    link VARCHAR(255),
    user_id VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);