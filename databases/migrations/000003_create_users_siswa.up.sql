CREATE TABLE IF NOT EXISTS siswa (
    -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    id INT AUTO_INCREMENT PRIMARY KEY,
    nis BIGINT UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    jenis_kelamin VARCHAR(10) NOT NULL,
    no_hp VARCHAR(20) NOT NULL,
    image_profile VARCHAR(255) NOT NULL,
    alamat TEXT NOT NULL,
    status_siswa_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);