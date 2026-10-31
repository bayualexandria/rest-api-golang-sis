CREATE TABLE pendaftaran (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nomor_pendaftaran VARCHAR(20) UNIQUE NOT NULL,
    id_calon_siswa BIGINT,
    id_jalur_pendaftaran BIGINT,
    tanggal_daftar DATETIME DEFAULT CURRENT_TIMESTAMP,
    status_verifikasi ENUM('Pending', 'Diverifikasi', 'Ditolak') DEFAULT 'Pending',
    catatan_verifikasi TEXT,
    status_kelulusan ENUM('Proses', 'Lulus', 'Tidak_lulus') DEFAULT 'Proses',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

     CONSTRAINT `fk_pendaftaran_id_calon_siswa`
        FOREIGN KEY (`id_calon_siswa`)
        REFERENCES `calon_siswa` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL,

    CONSTRAINT `fk_pendaftaran_id_jalur_pendaftaran`
        FOREIGN KEY (`id_jalur_pendaftaran`)
        REFERENCES `jalur_pendaftaran` (`id`)
         ON UPDATE CASCADE
        ON DELETE SET NULL
);