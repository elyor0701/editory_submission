DROP TYPE IF EXISTS "status_editor";
DROP TYPE IF EXISTS "status_reviewer";
DROP TYPE IF EXISTS "step";
DROP TYPE IF EXISTS "file_type";
DROP TYPE IF EXISTS "role";

DROP TABLE IF EXISTS "journal";
DROP TABLE IF EXISTS "article";
DROP TABLE IF EXISTS "draft";
DROP TABLE IF EXISTS "file";
DROP TABLE IF EXISTS "draft_reviewer";
DROP TABLE IF EXISTS "comment";
DROP TABLE IF EXISTS "coauthor";
DROP TABLE IF EXISTS "subject";
DROP TABLE IF EXISTS "journal_subject";
DROP TABLE IF EXISTS "country";
DROP TABLE IF EXISTS "city";
DROP TABLE IF EXISTS "university";
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "keyword";
DROP TABLE IF EXISTS "education";
DROP TABLE IF EXISTS "session";
DROP TABLE IF EXISTS "role";

CREATE TYPE "file_type" AS ENUM (
    'MANUSCRIPT',
    'COVER_LETTER',
    'SUPPLEMENTAL'
    );

CREATE TYPE "role_type" AS ENUM (
    'SUPERADMIN',
    'EDITOR',
    'PROOFREADER',
    'REVIEWER',
    'AUTHOR'
    );

CREATE TABLE "journal" (
                           "id" uuid PRIMARY KEY,
                           "cover_photo" varchar(255),
                           "title" varchar(255) not null,
                           "access" bool,
                           "description" text,
                           "price" int,
                           "isbn" varchar(255) not null unique,
                           "author" varchar(255),
                           "created_at" timestamp default CURRENT_TIMESTAMP,
                           "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "article" (
                           "id" uuid PRIMARY KEY,
                           "journal_id" uuid not null,
                           "article_type" varchar(255),
                           "title" varchar(255),
                           "author_id" uuid,
                           "description" text,
                           "created_at" timestamp default CURRENT_TIMESTAMP,
                           "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "file" (
                        "id" uuid PRIMARY KEY,
                        "url" varchar not null,
                        "file_type" file_type,
                        "draft_id" uuid,
                        "article_id" uuid
);

CREATE TABLE "coauthor" (
                            "id" uuid PRIMARY KEY,
                            "article_id" uuid,
                            "user_id" uuid
);

CREATE TABLE "subject" (
                           "id" uuid PRIMARY KEY,
                           "title" varchar(255)
);

CREATE TABLE "journal_subject" (
                                   "id" uuid PRIMARY KEY,
                                   "journal_id" uuid,
                                   "subject_id" uuid
);

CREATE TABLE "country" (
                           "id" uuid PRIMARY KEY,
                           "title" varchar(255)
);

CREATE TABLE "city" (
                        "id" uuid PRIMARY KEY,
                        "title" varchar(255),
                        "country_id" uuid
);

CREATE TABLE "university" (
                              "id" uuid PRIMARY KEY,
                              "title" varchar(255),
                              "logo" varchar(255)
);

CREATE TABLE "user" (
                        "id" uuid PRIMARY KEY,
                        "username" varchar(255),
                        "first_name" varchar(255),
                        "last_name" varchar(255),
                        "phone" varchar(255),
                        "extra_phone" varchar(255),
                        "email" varchar(255) not null unique,
                        "email_verification" bool default false,
                        "password" varchar(255),
                        "country_id" uuid,
                        "city_id" uuid,
                        "prof_sphere" varchar(255),
                        "degree" varchar(255),
                        "address" varchar(255),
                        "post_code" varchar(255),
                        "created_at" timestamp default CURRENT_TIMESTAMP,
                        "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "keyword" (
                           "id" uuid PRIMARY KEY,
                           "word" varchar(255),
                           "user_id" uuid
);

CREATE TABLE "education" (
                             "id" uuid PRIMARY KEY,
                             "user_id" uuid,
                             "university_id" uuid,
                             "start_year" timestamp,
                             "end_year" timestamp,
                             "still_studying" bool
);

CREATE TABLE "session" (
                           "id" uuid PRIMARY KEY,
                           "user_id" uuid,
                           "role_id" uuid,
                           "ip" inet,
                           "data" text,
                           "expires_at" timestamp,
                           "created_at" timestamp default CURRENT_TIMESTAMP,
                           "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "role" (
                        "id" uuid PRIMARY KEY,
                        "user_id" uuid,
                        "role_type" role_type,
                        "journal_id" uuid
);

ALTER TABLE "article" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");

ALTER TABLE "article" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");

ALTER TABLE "file" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id");

ALTER TABLE "coauthor" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id");

ALTER TABLE "coauthor" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "journal_subject" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");

ALTER TABLE "journal_subject" ADD FOREIGN KEY ("subject_id") REFERENCES "subject" ("id");

ALTER TABLE "city" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id");

ALTER TABLE "keyword" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "education" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "education" ADD FOREIGN KEY ("university_id") REFERENCES "university" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "role" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id");