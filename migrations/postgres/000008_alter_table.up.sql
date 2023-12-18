ALTER TABLE "email_verification" ADD COLUMN "user_id" UUID NOT NULL;

alter table "edition" add column title varchar(255);
alter table "edition" add column description varchar;

alter table "journal_data" add column "short_desc" varchar;