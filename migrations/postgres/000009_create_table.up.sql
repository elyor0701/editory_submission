-- create table "article" (
--     "id" uuid primary key,
--     "title" varchar,
--     "description" text,
--     "journal_id" uuid not null,
--     "edition" int not null,
--     "file" varchar,
--     "author" varchar,
--     "content" text,
--     "created_at" timestamp default CURRENT_TIMESTAMP,
--     "updated_at" timestamp default CURRENT_TIMESTAMP
-- );
--
-- create table "journal_author" (
--     "id" uuid primary key,
--     "journal_id" uuid not null,
--     "full_name" varchar,
--     "photo" varchar,
--     "email" varchar,
--     "university_id" uuid,
--     "faculty_id" uuid,
--     "created_at" timestamp default CURRENT_TIMESTAMP,
--     "updated_at" timestamp default CURRENT_TIMESTAMP
-- );
--
-- alter table "article" add foreign key ("journal_id") references "journal"("id") on delete set null ;
-- alter table "journal_author" add foreign key ("journal_id") references "journal"("id") on delete cascade ;
-- alter table "journal_author" add foreign key ("university_id") references "university"("id") on delete set null ;
-- alter table "user" add column "university_id" uuid;
-- alter table "user" add foreign key ("university_id") references "university"("id") on delete set null ;

create type "draft_step" as enum (
    'AUTHOR',
    'EDITOR',
    'REVIEWER'
);

create type "draft_editor_status" as enum (
    'NEW',
    'PENDING',
    'DONE'
);

create type "draft_reviewer_status" as enum (
    'NEW',
    'PENDING',
    'DONE'
);

alter type "article_status" add value 'BACK_FOR_CORRECTION';
alter type "reviewer_ans" add value 'BACK_FOR_CORRECTION';

alter table "draft" add column "step" draft_step default 'AUTHOR';
alter table "draft" add column "editor_status" draft_editor_status default 'NEW';
alter table "draft" add column "reviewer_status" draft_reviewer_status default 'NEW';
alter table "draft" add column "group_id" uuid unique;
alter table "draft" add column "manuscript" varchar;
alter table "draft" add column "cover_letter" varchar;
alter table "draft" add column "supplemental" varchar;
alter table "draft" add column "editor_manuscript_comment" text;
alter table "draft" add column "editor_cover_letter_comment" text;
alter table "draft" add column "editor_supplemental_comment" text;

alter table "article_reviewer" add column "manuscript_comment" text;
alter table "article_reviewer" add column "cover_letter_comment" text;
alter table "article_reviewer" add column "supplemental_comment" text;