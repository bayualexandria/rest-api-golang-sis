CREATE TABLE pendaftaran_nilai_calon_siswa (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_calon_siswa BIGINT,
    semester BIGINT NOT NULL, -- Semester 1-5
    mata_pelajaran VARCHAR(50) NOT NULL, -- Contoh: Matematika, Bahasa Indonesia, IPA
    nilai DECIMAL(5,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

     CONSTRAINT `fk_pendaftaran_nilai_calon_siswa`
        FOREIGN KEY (`id_calon_siswa`)
        REFERENCES `calon_siswa` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);