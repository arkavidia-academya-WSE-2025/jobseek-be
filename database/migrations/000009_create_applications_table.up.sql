CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE app_status AS ENUM ('pending', 'accepted', 'rejected');

CREATE TABLE applications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    full_name TEXT NOT NULL,
    address TEXT NOT NULL,
    application_status app_status DEFAULT 'pending',
    cv_path TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    job_id UUID NOT NULL REFERENCES jobs(id),
    job_seeker_id UUID NOT NULL REFERENCES users(id)
);
