-- Drop trigger from users table
DROP TRIGGER IF EXISTS trigger_update_user_updated_at ON users;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop pgcrypto extension
DROP EXTENSION IF EXISTS "pgcrypto";