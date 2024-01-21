alter table "draft" drop column "editor_id";
alter table "draft" drop column "editor_comment";
alter table "draft" drop column "manuscript";
alter table "draft" drop column "cover_letter";
alter table "draft" drop column "supplemental";
alter table "draft" drop column "editor_manuscript_comment";
alter table "draft" drop column "editor_cover_letter_comment";
alter table "draft" drop column "editor_supplemental_comment";

alter table "file" drop column "article_id";

alter table "article_reviewer" drop column "manuscript_comment";
alter table "article_reviewer" drop column "cover_letter_comment";
alter table "article_reviewer" drop column "supplemental_comment";
alter table "article_reviewer" rename to "draft_checker";

alter table "draft_checker" rename column "reviewer_id" to "checker_id";
alter table "draft_checker" rename column "article_id" to "draft_id";

alter type "article_status" add value 'DRAFT';

create type "checker_type" as enum(
    'EDITOR',
    'REVIEWER'
);

alter table "draft_checker" add column type checker_type;
alter table "draft_checker" drop column status;

drop type "reviewer_ans";

create type "checker_status" as enum (
    'NEW',
    'PENDING',
    'APPROVED',
    'REJECTED',
    'BACK_FOR_CORRECTION',
    'APPROVED_WITH_CORRECTION',
    'REJECTED_WITH_CORRECTION'
);

alter table "draft_checker" add column status checker_status default 'NEW';

create table "file_comment" (
  id uuid primary key,
  type file_type,
  file_id uuid,
  draft_checker_id uuid,
  comment text,
  created_at timestamp default CURRENT_TIMESTAMP,
  updated_at timestamp default CURRENT_TIMESTAMP
);

alter table "file_comment" add foreign key ("file_id") references file("id") on delete cascade;
alter table "file_comment" add foreign key ("draft_checker_id") references "draft_checker"("id") on delete set null;

alter table "draft" add column conflict bool default false;
alter table "draft" add column availability text;
alter table "draft" add column funding text;
alter table "draft" add column draft_step varchar;