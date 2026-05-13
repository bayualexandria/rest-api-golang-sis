CREATE TABLE kelas (
    -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan SERIAL atau BIGSERIAL untuk kolom id
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama_kelas VARCHAR(255) NOT NULL,
    jurusan VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);