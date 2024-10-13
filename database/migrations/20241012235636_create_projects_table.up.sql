CREATE SCHEMA IF NOT EXISTS bastion;

CREATE TABLE bastion.projects
(
    id           UUID         NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description  TEXT,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT projects_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE bastion.projects IS 'This table manages and organizes projects, allowing efficient tracking of project details, statuses, and associated resources within the authentication and authorization workflow.';

