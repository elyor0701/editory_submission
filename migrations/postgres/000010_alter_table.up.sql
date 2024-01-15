create table "article" (
    "id" uuid primary key,
    "title" varchar,
    "description" text,
    "journal_id" uuid not null,
    "edition" int not null,
    "file" varchar,
    "author" varchar,
    "content" text,
    "created_at" timestamp default CURRENT_TIMESTAMP,
    "updated_at" timestamp default CURRENT_TIMESTAMP
);

create table "journal_author" (
    "id" uuid primary key,
    "journal_id" uuid not null,
    "full_name" varchar,
    "photo" varchar,
    "email" varchar,
    "university_id" uuid,
    "faculty_id" uuid,
    "created_at" timestamp default CURRENT_TIMESTAMP,
    "updated_at" timestamp default CURRENT_TIMESTAMP
);

alter table "article" add foreign key ("journal_id") references "journal"("id") on delete set null ;
alter table "journal_author" add foreign key ("journal_id") references "journal"("id") on delete cascade ;
alter table "journal_author" add foreign key ("university_id") references "university"("id") on delete set null ;
alter table "user" add column "university_id" uuid;
alter table "user" add foreign key ("university_id") references "university"("id") on delete set null ;