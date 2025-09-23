-- Create the enum type for roles
CREATE TYPE conversation_role AS ENUM ('member', 'admin', 'moderator');

-- Create conversations_members table and connect it to conversation id and user id
CREATE TABLE conversations_members(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    conversation_id UUID REFERENCES conversations(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role conversation_role NOT NULL DEFAULT 'member',
    is_muted BOOLEAN NOT NULL DEFAULT FALSE,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    muted_until TIMESTAMPTZ,
    UNIQUE(conversation_id, user_id)
);