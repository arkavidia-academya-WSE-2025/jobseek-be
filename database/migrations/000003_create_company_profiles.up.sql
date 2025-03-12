CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table company_profiles(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
    user_id UUID NOT NULL REFERENCES users(id) on delete CASCADE,
    photo_url text,
    description text,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)