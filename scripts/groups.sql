CREATE TABLE groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    info JSON NOT NULL DEFAULT '{}',
    is_active VARCHAR(10) DEFAULT 'inactive' CHECK (is_active IN ('active', 'inactive')),
    is_deleted TIMESTAMP NULL,
    picture INTEGER,
    members_count INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_groups_organization_id ON groups(organization_id);
CREATE INDEX idx_groups_is_active ON groups(is_active);
CREATE INDEX idx_groups_is_deleted ON groups(is_deleted);
