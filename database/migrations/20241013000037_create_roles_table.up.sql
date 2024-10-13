CREATE TABLE bastion.roles
(
    id           UUID        NOT NULL,
    tenant_id    UUID,
    display_name VARCHAR(100),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT roles_id_pk PRIMARY KEY (id),
    CONSTRAINT roles_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES bastion.tenants (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.roles IS 'This table defines various roles that can be assigned to users, allowing the implementation of role-based access control within the application.';

