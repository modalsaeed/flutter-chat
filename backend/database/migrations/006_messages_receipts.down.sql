-- Drop trigger from messages_receipts table
DROP TRIGGER IF EXISTS trigger_update_messages_receipts_updated_at ON messages_receipts;

-- Drop messages_receipts table
DROP TABLE IF EXISTS messages_receipts;

-- Drop receipt_status ENUM type
DROP TYPE IF EXISTS receipt_status