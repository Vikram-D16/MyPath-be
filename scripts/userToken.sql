
CREATE TABLE user_tokens (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    token TEXT NOT NULL UNIQUE,
    is_active BOOL not null default TRUE,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);