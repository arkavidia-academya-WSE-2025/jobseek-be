CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
create table jobs(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  requirements TEXT NOT NULL,
  location TEXT NOT NULL,
  salary INT,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
)