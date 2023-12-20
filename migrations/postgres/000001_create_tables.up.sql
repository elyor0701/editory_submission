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

CREATE TYPE "journal_status" AS ENUM (
    'ACTIVE',
    'INACTIVE'
    );

CREATE TYPE "gender" AS ENUM (
    'MALE',
    'FEMALE'
    );

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

CREATE TABLE "journal" (
                           "id" uuid PRIMARY KEY,
                           "cover_photo" varchar(255),
                           "title" varchar(255) not null,
                           "access" bool,
                           "description" text,
                           "price" int,
                           "isbn" varchar(255) not null unique,
                           "author" varchar(255),
                            "status" journal_status default 'ACTIVE' not null,
                           "acceptance_rate" varchar,
                            "submission_to_final_decision" varchar,
                           "acceptance_to_publication" varchar,
                           "citation_indicator" varchar,
                           "impact_factor" varchar,
                           "short_description" varchar,
                           "created_at" timestamp default CURRENT_TIMESTAMP,
                           "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "journal_data" (
                                journal_id UUID NOT NULL,
                                short_desc varchar,
                                text TEXT,
                                type "journal_data_type" NOT NULL,
                                UNIQUE ("journal_id", "type")
);

CREATE TABLE "article" (
                           "id" uuid PRIMARY KEY,
                           "journal_id" uuid not null,
                           "type" varchar(255),
                           "title" varchar(255),
                           "author_id" uuid,
                           "description" text,
                           "created_at" timestamp default CURRENT_TIMESTAMP,
                           "updated_at" timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE "edition" (
                           id UUID PRIMARY KEY,
                           edition INT NOT NULL,
                           file VARCHAR,
                           journal_id UUID NOT NULL,
                            title varchar(255),
                            description varchar,
                           created_at TIMESTAMP DEFAULT current_timestamp,
                           updated_at TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE "file" (
                        "id" uuid PRIMARY KEY,
                        "url" varchar not null,
                        "type" file_type,
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
                           "title" varchar(255),
                            "title_ru" varchar,
                            "title_uz" varchar
);

CREATE TABLE "city" (
                        "id" uuid PRIMARY KEY,
                        "title" varchar(255),
                        "country_id" uuid,
                        "title_ru" varchar,
                        "title_uz" varchar
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
                        "gender" gender,
                        "is_completed" boolean,
                        "created_at" timestamp default CURRENT_TIMESTAMP,
                        "updated_at" timestamp default CURRENT_TIMESTAMP
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

create table email_verification (
                        email varchar(255) not null,
                        token varchar not null,
                        sent bool default false,
                        expires_at timestamp default CURRENT_TIMESTAMP + INTERVAL '1 day',
                        created_at timestamp default CURRENT_TIMESTAMP,
                        "user_id" UUID not null,
                        UNIQUE(email, token)
);

ALTER TABLE "article" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id") ON DELETE SET NULL;

ALTER TABLE "article" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id") ON DELETE SET NULL;

ALTER TABLE "file" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id") ON DELETE CASCADE;

ALTER TABLE "coauthor" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id") ON DELETE CASCADE;

ALTER TABLE "coauthor" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE;

ALTER TABLE "journal_subject" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id") ON DELETE CASCADE ;

ALTER TABLE "journal_subject" ADD FOREIGN KEY ("subject_id") REFERENCES "subject" ("id") ON DELETE CASCADE ;

ALTER TABLE "city" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id") ON DELETE SET NULL ;

ALTER TABLE "user" ADD FOREIGN KEY ("country_id") REFERENCES "country" ("id") ON DELETE SET NULL;

ALTER TABLE "user" ADD FOREIGN KEY ("city_id") REFERENCES "city" ("id") ON DELETE SET NULL;

ALTER TABLE "education" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ;

ALTER TABLE "education" ADD FOREIGN KEY ("university_id") REFERENCES "university" ("id") ON DELETE CASCADE ;

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ;

ALTER TABLE "session" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id") ON DELETE CASCADE ;

ALTER TABLE "role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON DELETE CASCADE ;

ALTER TABLE "role" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id") ON DELETE CASCADE ;

ALTER TABLE "journal_data" ADD FOREIGN KEY ("journal_id") REFERENCES "journal" ("id") on delete cascade ;

ALTER TABLE edition ADD FOREIGN KEY ("journal_id") REFERENCES journal("id") on delete set null ;