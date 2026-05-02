CREATE TABLE IF NOT EXISTS guru(
    -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan SERIAL atau BIGSERIAL untuk kolom id
    id BIGSERIAL PRIMARY KEY,
    nip BIGINT UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    jenis_kelamin VARCHAR(10) NOT NULL,
    no_hp VARCHAR(20) NOT NULL,
    image_profile VARCHAR(255) NOT NULL,
    alamat TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
)