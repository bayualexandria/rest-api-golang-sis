CREATE TABLE IF NOT EXISTS siswa (
    id BIGINT PRIMARY KEY,
    nis BIGINT UNIQUE NOT NULL,
    nama VARCHAR(255) NOT NULL,
    jenis_kelamin VARCHAR(10) NOT NULL,
    no_hp VARCHAR(20) NOT NULL,
    image_profile VARCHAR(255) NOT NULL,
    alamat TEXT NOT NULL,
    status_siswa_id BIGINT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL
);