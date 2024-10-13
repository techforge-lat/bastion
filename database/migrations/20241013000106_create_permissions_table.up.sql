CREATE TABLE bastion.permissions
(
    id                       UUID        NOT NULL,
    role_id                  UUID        NOT NULL,
    project_module_id        UUID, -- permission can be granted for all actions in a project module
    project_module_action_id UUID, -- or permissions can be defined per action
    created_at               TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at               TIMESTAMPTZ,
    CONSTRAINT permissions_id_pk PRIMARY KEY (id),
    CONSTRAINT permissions_role_id_fk FOREIGN KEY (role_id) REFERENCES bastion.roles (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT permissions_project_module_id_fk FOREIGN KEY (project_module_id) REFERENCES bastion.project_modules (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT permissions_project_module_action_id_fk FOREIGN KEY (project_module_action_id) REFERENCES bastion.project_module_actions (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.permissions IS 'This table manages permissions assigned to roles, allowing granular access control to project modules and their associated actions within the application.';

