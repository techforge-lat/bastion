CREATE TABLE bastion.user_roles
(
    id         UUID        NOT NULL,
    tenant_id  UUID,
    user_id    UUID        NOT NULL,
    role_id    UUID        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ,
    CONSTRAINT user_roles_id_pk PRIMARY KEY (id),
    CONSTRAINT user_roles_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES bastion.tenants (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT user_roles_user_id_fk FOREIGN KEY (user_id) REFERENCES bastion.users (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT user_roles_role_id_fk FOREIGN KEY (role_id) REFERENCES bastion.roles (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.user_roles IS 'This table associates users with their assigned roles within a tenant, facilitating role-based access control and user permission management within the application.';

