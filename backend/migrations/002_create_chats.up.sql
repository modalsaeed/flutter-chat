-- Create chats table with auto-generated UUID
CREATE TABLE chats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50),
    is_group BOOLEAN NOT NULL DEFAULT FALSE,
    created_by UUID REFERENCES users(id),
    description VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    avatar_url TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Attach trigger to chats table
CREATE TRIGGER trigger_update_chats_updated_at
BEFORE UPDATE ON chats
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

