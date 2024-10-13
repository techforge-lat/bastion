CREATE TABLE bastion.users
(
    id           UUID         NOT NULL,
    tenant_id    UUID,
    display_name VARCHAR(100),
    email        VARCHAR(256) NOT NULL,
    password     VARCHAR(256) NOT NULL,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT users_id_pk PRIMARY KEY (id),
    CONSTRAINT users_email_uk UNIQUE (email),
    CONSTRAINT users_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES bastion.tenants (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.users IS 'This table stores user account information, including authentication details, user identifiers, and associations with tenants for access control.';

