CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE Table jobseeker_skills(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4 (),
  jobseeker_id UUID NOT NULL REFERENCES jobseeker_profiles(id)
)