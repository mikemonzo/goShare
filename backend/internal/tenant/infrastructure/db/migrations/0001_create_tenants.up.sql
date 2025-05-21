CREATE TABLE IF NOT EXISTS tenants (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    branding TEXT,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL
);
