CREATE TABLE linked_social_accounts (
    -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(username),
    provider_name VARCHAR(255) NOT NULL,
    provider_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)