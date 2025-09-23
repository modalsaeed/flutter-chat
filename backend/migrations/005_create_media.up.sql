-- Create enum for media types
CREATE TYPE media_type AS ENUM ('image', 'video', 'audio', 'file');

-- Create media table
CREATE TABLE media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uploader_id UUID REFERENCES users(id) ON DELETE SET NULL,
    chat_id UUID REFERENCES chats(id) ON DELETE CASCADE,
    message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    type media_type NOT NULL,
    mime_type TEXT,
    size_bytes BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)