CREATE TABLE personal_access_tokens (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  tokenable_type varchar(255) not null,
  tokenable_id BIGINT not null,
  name varchar(255) not null,
  token varchar(255) not null,
  abilities text,
  last_used_at TIMESTAMP NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
)