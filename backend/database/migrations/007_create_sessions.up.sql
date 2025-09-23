-- Create sessions table
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    ip_address INET,          
    user_agent TEXT,          
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMPTZ NOT NULL,  
    revoked BOOLEAN NOT NULL DEFAULT FALSE  
);