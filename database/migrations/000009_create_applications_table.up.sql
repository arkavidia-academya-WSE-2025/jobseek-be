CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'app_status') THEN
        CREATE TYPE app_status AS ENUM ('pending', 'accepted', 'rejected');
    END IF;
END
$$;

CREATE TABLE IF NOT EXISTS applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    address TEXT NOT NULL,
    application_status app_status DEFAULT 'pending',
    cv_path TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    job_id UUID NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
    job_seeker_id UUID NOT NULL REFERENCES users(id)
);
