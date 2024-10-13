CREATE TABLE bastion.tenant_project_modules
(
    id                UUID        NOT NULL,
    tenant_id         UUID        NOT NULL,
    project_module_id UUID        NOT NULL,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at        TIMESTAMPTZ,
    CONSTRAINT tenant_project_modules_id_pk PRIMARY KEY (id),
    CONSTRAINT tenant_project_modules_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES bastion.tenants (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT tenant_project_modules_project_module_id_fk FOREIGN KEY (project_module_id) REFERENCES bastion.project_modules (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.tenant_project_modules IS 'This table establishes a relationship between tenants and the project modules they are associated with, facilitating the management of module availability per tenant.';

