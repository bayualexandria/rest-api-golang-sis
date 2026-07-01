CREATE TABLE password_reset_tokens (
      -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);