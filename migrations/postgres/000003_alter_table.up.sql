ALTER TABLE edition ADD CONSTRAINT edition_unique_edition_journal_id UNIQUE ("edition", "journal_id");

ALTER TABLE "journal_subject" ADD CONSTRAINT journal_subject_unique_journal_subject UNIQUE ("journal_id", subject_id);

ALTER TABLE "role" ADD CONSTRAINT role_user_journal_role UNIQUE (user_id, journal_id, role_type);