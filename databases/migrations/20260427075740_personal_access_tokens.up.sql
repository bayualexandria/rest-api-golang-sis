CREATE TABLE personal_access_tokens (
  id BIGINT PRIMARY KEY,
  tokenable_type varchar(255) not null,
  tokenable_id BIGINT not null,
  name varchar(255) not null,
  token varchar(255) not null,
  abilities text,
  last_used_at datetime,
  expires_at datetime,
  created_at datetime,
  updated_at datetime
)