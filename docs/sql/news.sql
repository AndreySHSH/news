-- =============================================================================
-- Diagram Name: project
-- Created on: 11.04.2022 13:47:57
-- Diagram Version:
-- =============================================================================


CREATE TABLE "news" (
    "newsId" BIGSERIAL NOT NULL,
    "title" text NOT NULL,
    "content" text NOT NULL,
    "tagIds" int8[],
    "categoryId" int8 NOT NULL,
    "createdAt" timestamp with time zone NOT NULL DEFAULT now(),
    "statusId" int4 NOT NULL,
    PRIMARY KEY("newsId")
);

CREATE TABLE "categories" (
    "categoryId" BIGSERIAL NOT NULL,
    "title" text NOT NULL,
    "statusId" int4 NOT NULL,
    PRIMARY KEY("categoryId")
);

CREATE TABLE "tags" (
    "tagId" BIGSERIAL NOT NULL,
    "title" text NOT NULL,
    "statusId" int4 NOT NULL,
    PRIMARY KEY("tagId")
);

CREATE TABLE "statuses" (
    "statusId" int4 NOT NULL,
    PRIMARY KEY("statusId")
);


ALTER TABLE "news" ADD CONSTRAINT "Ref_news_to_categories" FOREIGN KEY ("categoryId")
    REFERENCES "categories"("categoryId")
    MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "news" ADD CONSTRAINT "Ref_news_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
    MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "categories" ADD CONSTRAINT "Ref_categories_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
    MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;

ALTER TABLE "tags" ADD CONSTRAINT "Ref_tags_to_statuses" FOREIGN KEY ("statusId")
    REFERENCES "statuses"("statusId")
    MATCH SIMPLE
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    NOT DEFERRABLE;


