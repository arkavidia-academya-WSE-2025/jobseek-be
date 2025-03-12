CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE jobseeker_profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id UUID  NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    photo_url TEXT,
    headline TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)