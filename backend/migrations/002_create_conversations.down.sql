-- Drop trigger from conversations table
DROP TRIGGER IF EXISTS trigger_update_conversations_updated_at on conversations;

-- Drop conversations table
DROP TABLE IF EXISTS conversations;