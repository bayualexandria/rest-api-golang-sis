CREATE TABLE linked_social_accounts (
    -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- id INT AUTO_INCREMENT PRIMARY KEY,
    -- Jika menggunakan PostgreSQL, gunakan SERIAL
    -- id SERIAL PRIMARY KEY
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    provider_name VARCHAR(255) NOT NULL,
    provider_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)