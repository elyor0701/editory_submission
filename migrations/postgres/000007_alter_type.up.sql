ALTER TYPE journal_data_type ADD VALUE 'JOURNAL_PROFILE' AFTER 'ARTICLE_PROCESSING_CHARGES';

alter table "journal" drop column "author";