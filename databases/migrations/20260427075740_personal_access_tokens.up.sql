CREATE TABLE personal_access_tokens (
  -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan SERIAL atau BIGSERIAL untuk kolom id
  id BIGSERIAL PRIMARY KEY,
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