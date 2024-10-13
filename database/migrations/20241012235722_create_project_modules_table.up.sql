CREATE TABLE bastion.project_modules
(
    id           UUID         NOT NULL,
    project_id   UUID         NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    is_active    BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT project_modules_id_pk PRIMARY KEY (id),
    CONSTRAINT project_modules_project_id_fk FOREIGN KEY (project_id) REFERENCES bastion.projects (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.project_modules IS 'This table contains the modules associated with each project, defining features or components that contribute to the project’s functionality and its active status.';

