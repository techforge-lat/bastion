CREATE TABLE bastion.tenants
(
    id           UUID         NOT NULL,
    project_id   UUID         NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT tenants_id_pk PRIMARY KEY (id),
    CONSTRAINT tenants_project_id_fk FOREIGN KEY (project_id) REFERENCES bastion.projects (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.tenants IS 'This table manages different tenants (clients or organizations) using the application, linking each tenant to a specific project and providing essential metadata.';

