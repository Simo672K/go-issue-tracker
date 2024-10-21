-- Create the database if it does not exist
CREATE DATABASE IF NOT EXISTS issue_tracker;

-- Connect to the issue_tracker database
\c issue_tracker;

-- Enable the uuid-ossp extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create ENUM for issue statuses
CREATE TYPE i_status AS ENUM ('IN_PROGRESS', 'OPEN', 'CLOSED');

-- Create the user table with UUID primary key
CREATE TABLE "user" (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the profile table with foreign key to user
CREATE TABLE profile (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  user_id UUID UNIQUE NOT NULL,
  username VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE
);

-- Create the project table
CREATE TABLE project (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  project_name VARCHAR(255) NOT NULL,
  project_progress FLOAT DEFAULT 0 NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the project_owner table with CASCADE delete on both foreign keys
CREATE TABLE project_owner (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  owner_id UUID NOT NULL,
  project_id UUID UNIQUE NOT NULL,
  FOREIGN KEY (owner_id) REFERENCES profile (id) ON DELETE CASCADE,
  FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);

-- Create the project_manager table
CREATE TABLE project_manager (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  profile_id UUID NOT NULL,
  project_id UUID NOT NULL,
  FOREIGN KEY (profile_id) REFERENCES profile (id) ON DELETE CASCADE,
  FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);

-- Create the project_dev table
CREATE TABLE project_dev (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  profile_id UUID NOT NULL,
  project_id UUID NOT NULL,
  FOREIGN KEY (profile_id) REFERENCES profile (id) ON DELETE CASCADE,
  FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);

-- Create the issue table with ENUM status and CASCADE on project_id delete
CREATE TABLE issue (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
  title VARCHAR(255) NOT NULL,
  project_id UUID,
  issue_status i_status DEFAULT 'OPEN' NOT NULL, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
);

-- Create the issue_assigned table with CASCADE delete for all foreign keys
CREATE TABLE issue_assigned (
  issue_id UUID UNIQUE NOT NULL,
  assigned_by UUID NOT NULL,
  assigned_to UUID NOT NULL,
  FOREIGN KEY (issue_id) REFERENCES issue (id) ON DELETE CASCADE,
  FOREIGN KEY (assigned_by) REFERENCES profile (id) ON DELETE CASCADE,
  FOREIGN KEY (assigned_to) REFERENCES profile (id) ON DELETE CASCADE
);

-- Create indexes for performance optimization
CREATE INDEX idx_profile_user_id ON profile (user_id);
CREATE INDEX idx_project_owner_id ON project_owner (owner_id);
CREATE INDEX idx_project_owner_project_id ON project_owner (project_id);
CREATE INDEX idx_project_manager_profile_id ON project_manager (profile_id);
CREATE INDEX idx_project_manager_project_id ON project_manager (project_id);
CREATE INDEX idx_project_dev_profile_id ON project_dev (profile_id);
CREATE INDEX idx_project_dev_project_id ON project_dev (project_id);
CREATE INDEX idx_issue_project_id ON issue (project_id);
CREATE INDEX idx_issue_assigned_issue_id ON issue_assigned (issue_id);
CREATE INDEX idx_issue_assigned_assigned_by ON issue_assigned (assigned_by);
CREATE INDEX idx_issue_assigned_assigned_to ON issue_assigned (assigned_to);
