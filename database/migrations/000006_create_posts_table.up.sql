create table posts(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  user_id UUID NOT NULL REFERENCES users(id) on delete CASCADE,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
)