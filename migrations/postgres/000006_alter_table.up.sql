ALTER TABLE "journal" ADD COLUMN author_id UUID;
ALTER TABLE "journal" ADD FOREIGN KEY ("author_id") REFERENCES "user"("id") ON DELETE SET NULL;