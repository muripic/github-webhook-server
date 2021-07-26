CREATE SCHEMA github;

CREATE TABLE github.issue (
    "id" bigint PRIMARY KEY,
    "title" text NOT NULL,
    "number" integer NOT NULL,
    "state" character varying(10) NOT NULL,
    "body" text,
    "created_by" character varying(64) NOT NULL,
    "url" character varying(255) NOT NULL,
    "created_at" date NOT NULL,
    "updated_at" date,
    "closed_at" date
);

CREATE TABLE github.issue_comment (
    "id" bigint PRIMARY KEY,
    "issue_id" integer,
    "body" text,
    "created_by" character varying(64) NOT NULL,
    "created_at" date NOT NULL,
    "updated_at" date,
    "url" character varying(255) UNIQUE NOT NULL,
    CONSTRAINT fk_issue FOREIGN KEY(issue_id) REFERENCES github.issue(id)
);

CREATE TABLE github.label (
    "id" bigint NOT NULL PRIMARY KEY,
    "label" character varying(64) UNIQUE NOT NULL
);

CREATE TABLE github.issue_label (
    "label_id" bigint NOT NULL,
    "issue_id" bigint NOT NULL,
    PRIMARY KEY(label_id, issue_id),
    CONSTRAINT fk_issue FOREIGN KEY(issue_id) REFERENCES github.issue(id),
    CONSTRAINT fk_label FOREIGN KEY(label_id) REFERENCES github.label(id)
)
