alter table "edition" drop constraint edition_unique_edition_journal_id;
alter table "journal_subject" drop constraint journal_subject_unique_journal_subject;
alter table "role" drop constraint role_user_journal_role;