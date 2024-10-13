CREATE TABLE bastion.tenant_sign_in_providers
(
    id                        UUID        NOT NULL,
    tenant_id                 UUID        NOT NULL,
    sign_in_provider_id       UUID        NOT NULL,
    is_active                 BOOLEAN     NOT NULL DEFAULT FALSE,
    override_sign_in_metadata BOOLEAN     NOT NULL DEFAULT FALSE,
    metadata                  JSONB       NOT NULL DEFAULT '{}',
    created_at                TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at                TIMESTAMPTZ,
    CONSTRAINT tenant_sign_in_providers_id_pk PRIMARY KEY (id),
    CONSTRAINT tenant_sign_in_providers_tenant_id_fk FOREIGN KEY (tenant_id) REFERENCES bastion.tenants (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
    CONSTRAINT tenant_sign_in_providers_sign_in_provider_id_fk FOREIGN KEY (sign_in_provider_id) REFERENCES bastion.sign_in_providers (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

COMMENT ON TABLE bastion.tenant_sign_in_providers IS 'This table links tenants to their configured sign-in providers, allowing customization of authentication options and management of provider metadata specific to each tenant.';

