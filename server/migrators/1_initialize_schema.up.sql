-- set timezone
SET TIMEZONE="Asia/Dhaka";

-- create users table
CREATE TABLE IF NOT EXISTS "users"(
id SERIAL PRIMARY KEY,
identifier VARCHAR(50) UNIQUE NOT NULL,
password VARCHAR(200) NOT NULL,
first_name VARCHAR(50),
last_name VARCHAR(50),
created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
updated_at TIMESTAMP DEFAULT current_timestamp
);

-- create table attendace
CREATE TABLE IF NOT EXISTS "events"(
  id SERIAL PRIMARY KEY,
  hash TEXT,
  type_tx VARCHAR,
  employee_id NUMERIC,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP DEFAULT current_timestamp
)
