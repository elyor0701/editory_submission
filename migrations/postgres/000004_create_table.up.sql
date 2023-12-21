-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "keyword" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  word VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE "user_keyword" (
    user_id UUID,
    keyword_id UUID
);

CREATE TYPE email_template_type AS ENUM (
    'REGISTRATION',
    'RESET_PASSWORD',
    'ACCOUNT_DEACTIVATION',
    'NEW_ARTICLE_TO_REVIEW',
    'NEW_JOURNAL_USER'
);

CREATE TABLE "email_template" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255),
    description VARCHAR,
    type email_template_type UNIQUE NOT NULL,
    text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE notification_status AS ENUM (
    'NEW',
    'PENDING',
    'SENT',
    'FAILED'
);

CREATE TABLE "notification" (
    id UUID PRIMARY KEY,
    subject TEXT DEFAULT '',
    text TEXT DEFAULT '',
    email VARCHAR NOT NULL,
    status notification_status,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE "user_keyword" ADD FOREIGN KEY ("user_id") REFERENCES "user"("id") ON DELETE CASCADE;
ALTER TABLE "user_keyword" ADD FOREIGN KEY ("keyword_id") REFERENCES "keyword"("id") ON DELETE CASCADE;