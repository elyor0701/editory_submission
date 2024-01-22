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
alter type "article_status" add value 'CORRECTED';
alter type "reviewer_ans" add value 'BACK_FOR_CORRECTION';

alter table "draft" add column "step" draft_step default 'AUTHOR';
alter table "draft" add column "editor_status" draft_editor_status default 'NEW';
alter table "draft" add column "reviewer_status" draft_reviewer_status default 'NEW';
alter table "draft" add column "group_id" uuid;
alter table "draft" add column "manuscript" varchar;
alter table "draft" add column "cover_letter" varchar;
alter table "draft" add column "supplemental" varchar;
alter table "draft" add column "editor_manuscript_comment" text;
alter table "draft" add column "editor_cover_letter_comment" text;
alter table "draft" add column "editor_supplemental_comment" text;

alter table "article_reviewer" add column "manuscript_comment" text;
alter table "article_reviewer" add column "cover_letter_comment" text;
alter table "article_reviewer" add column "supplemental_comment" text;