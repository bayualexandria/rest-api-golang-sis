CREATE TABLE IF NOT EXISTS status_user (
    -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan SERIAL atau BIGSERIAL untuk kolom id
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);