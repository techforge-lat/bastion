CREATE TABLE bastion.sign_in_providers
(
    id           UUID         NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description  TEXT,
    icon         VARCHAR(100) NOT NULL,
    is_active    BOOLEAN      NOT NULL,
    metadata     JSONB        NOT NULL DEFAULT '{}',
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ,
    CONSTRAINT sign_in_providers_id_pk PRIMARY KEY (id)
);

COMMENT ON TABLE bastion.sign_in_providers IS 'This table manages various sign-in providers (e.g., Google, Facebook) available for user authentication, including their display information and status.';

