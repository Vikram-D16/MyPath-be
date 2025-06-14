CREATE TABLE organization_members (
    organization_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    PRIMARY KEY (organization_id, user_id)
);

CREATE INDEX idx_org_members_not_deleted ON organization_members(organization_id, user_id) WHERE deleted_at IS NULL;
