CREATE TABLE profile_sekolah(
        -- Jika menggunakan MySQL, gunakan AUTO_INCREMENT untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan SERIAL atau BIGSERIAL untuk kolom id
    id BIGSERIAL PRIMARY KEY,
    nama_sekolah VARCHAR(255) NOT NULL,
    alamat_sekolah TEXT NOT NULL,
    telepon_sekolah VARCHAR(20) NOT NULL,
    akreditasi VARCHAR(10) NOT NULL,
    image_profile VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
)