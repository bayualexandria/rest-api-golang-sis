CREATE TABLE dokumen_persyaratan_calon_siswa (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_calon_siswa BIGINT,
    jenis_dokumen ENUM('Ijazah', 'Kartu Keluarga', 'Akta Lahir', 'Pas Foto', 'Sertifikat') NOT NULL,
    nama_file VARCHAR(255) NOT NULL,
    path_file VARCHAR(255) NOT NULL,
    diunggah_pada TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

     CONSTRAINT `fk_dokumen_persyaratan_calon_siswa`
        FOREIGN KEY (`id_calon_siswa`)
        REFERENCES `calon_siswa` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);