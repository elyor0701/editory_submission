DROP TYPE IF EXISTS "article_status";
DROP TYPE IF EXISTS "reviewer_ans";

ALTER TABLE "article" DROP COLUMN "status";
ALTER TABLE "article" DROP COLUMN "editor_comment";
ALTER TABLE "article" DROP COLUMN "editor_id";

drop table "article_reviewer";