-- Create enum for receipt status
CREATE TYPE receipt_status AS ENUM ('sent', 'delivered', 'read');

-- Create messages receipts table
CREATE TABLE messages_receipts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message_id UUID REFERENCES messages(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    status receipt_status NOT NULL DEFAULT 'sent',
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (message_id, user_id)
)

-- Attach trigger to messages_receipts table
CREATE TRIGGER trigger_update_messages_receipts_updated_at
BEFORE UPDATE ON messages_receipts
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();