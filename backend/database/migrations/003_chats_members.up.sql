-- Create the enum type for roles
CREATE TYPE chat_role AS ENUM ('member', 'admin', 'moderator');

-- Create chats_members table and connect it to chat id and user id
CREATE TABLE chats_members(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chat_id UUID REFERENCES chats(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    role chat_role NOT NULL DEFAULT 'member',
    is_muted BOOLEAN NOT NULL DEFAULT FALSE,
    is_banned BOOLEAN NOT NULL DEFAULT FALSE,
    muted_until TIMESTAMPTZ,
    UNIQUE(chat_id, user_id)
);