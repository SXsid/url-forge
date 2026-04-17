--+goose up 
--+goose StatementBegin
CREATE TABLE url(
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(11) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd
