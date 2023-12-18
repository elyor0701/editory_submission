CREATE TABLE "edition" (
    id UUID PRIMARY KEY,
    edition INT NOT NULL,
    file VARCHAR,
    journal_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE edition ADD FOREIGN KEY ("journal_id") REFERENCES journal("id");