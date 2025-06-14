CREATE TYPE organization_status AS ENUM ('active', 'inactive');

CREATE TABLE organizations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_id INTEGER,
    title VARCHAR NOT NULL,
    description TEXT NOT NULL,
    address VARCHAR,
    country VARCHAR NOT NULL,
    state VARCHAR NOT NULL,
    city VARCHAR NOT NULL,
    zipcode VARCHAR(255),
    email VARCHAR,
    website VARCHAR,
    phone VARCHAR,
    about TEXT,
    is_active organization_status DEFAULT 'inactive',
    is_deleted TIMESTAMP,
    picture INTEGER,
    followers_count INTEGER DEFAULT 0,
    members_count INTEGER DEFAULT 0,
    settings JSONB DEFAULT '{}'::jsonb,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_organizations_title ON organizations(title);

CREATE INDEX idx_organizations_active ON organizations(is_active) WHERE is_deleted IS NULL;
