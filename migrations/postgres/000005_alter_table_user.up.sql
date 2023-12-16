CREATE TYPE "gender" AS ENUM (
    'MALE',
    'FEMALE'
);

ALTER TABLE "user" ADD COLUMN gender gender;

ALTER TABLE journal ADD COLUMN "acceptance_rate" VARCHAR;
ALTER TABLE journal ADD COLUMN "submission_to_final_decision" VARCHAR;
ALTER TABLE journal ADD COLUMN "acceptance_to_publication" VARCHAR;
ALTER TABLE journal ADD COLUMN "citation_indicator" VARCHAR;
ALTER TABLE journal ADD COLUMN "impact_factor" VARCHAR;

ALTER TABLE journal ADD COLUMN "short_description" VARCHAR;

CREATE TYPE "journal_data_type" AS ENUM (
    'EDITOR_SPOTLIGHT',
    'SPECIAL_ISSUE',
    'ABOUT_JOURNAL',
    'EDITORIAL_BARD',
    'PEER_REVIEW_PROCESS',
    'PUBLICATION_ETHICS',
    'ABSTRACTING_INDEXING',
    'ARTICLE_PROCESSING_CHARGES'
);

CREATE TABLE "journal_data" (
    journal_id UUID NOT NULL,
    text TEXT,
    type "journal_data_type" NOT NULL
);

ALTER TABLE "journal_data" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");
ALTER TABLE "journal_data" ADD CONSTRAINT "unique_journal_id_type" UNIQUE ("journal_id", "type");

