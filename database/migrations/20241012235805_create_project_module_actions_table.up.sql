CREATE TABLE bastion.project_module_actions
(
    id                UUID         NOT NULL,
    project_module_id UUID         NOT NULL,
    display_name      VARCHAR(250) NOT NULL,
    code              VARCHAR(250) NOT NULL,
    description       TEXT,
    is_active         BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at        TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ,
    CONSTRAINT project_module_actions_id_pk PRIMARY KEY (id),
    CONSTRAINT project_module_actions_project_module_id_fk FOREIGN KEY (project_module_id) REFERENCES bastion.project_modules (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.project_module_actions IS 'This table defines specific actions associated with each project module, including its active status and related metadata.';

