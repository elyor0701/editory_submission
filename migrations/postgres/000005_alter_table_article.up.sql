DROP TYPE IF EXISTS "article_status";
DROP TYPE IF EXISTS "reviewer_ans";

CREATE TYPE "article_status" AS ENUM (
    'NEW',
    'PENDING',
    'DENIED',
    'CONFIRMED',
    'PUBLISHED'
);

CREATE TYPE "reviewer_ans" AS ENUM (
    'PENDING',
    'APPROVED',
    'REJECTED'
);

ALTER TABLE "article" ADD COLUMN status article_status NOT NULL DEFAULT 'NEW';
ALTER TABLE "article" ADD COLUMN editor_comment TEXT;
ALTER TABLE "article" ADD COLUMN editor_id UUID;

CREATE TABLE "article_reviewer" (
    id UUID PRIMARY KEY ,
    reviewer_id UUID NOT NULL,
    article_id UUID NOT NULL,
    status reviewer_ans DEFAULT 'PENDING',
    comment TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE "article_reviewer" ADD FOREIGN KEY ("reviewer_id") REFERENCES "user"("id") ON DELETE SET NULL;
ALTER TABLE "article_reviewer" ADD FOREIGN KEY ("article_id") REFERENCES "article"("id") ON DELETE CASCADE;