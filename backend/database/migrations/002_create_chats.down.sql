-- Drop trigger from chats table
DROP TRIGGER IF EXISTS trigger_update_chats_updated_at ON chats;

-- Drop chats table
DROP TABLE IF EXISTS chats;