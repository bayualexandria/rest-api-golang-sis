CREATE TABLE jalur_pendaftaran (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nama_jalur VARCHAR(50) NOT NULL, -- Contoh: Zonasi, Prestasi Academic, Afirmasi
    deskripsi TEXT,
    kuota BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);