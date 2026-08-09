CREATE TABLE `nilai_siswa` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `penilaian_id` BIGINT(20) ,
    `siswa_id` BIGINT(20) ,
    `nilai` DECIMAL(5,2),
    `predikat` VARCHAR(10),
    `catatan` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_nilai_penilaian`
        FOREIGN KEY (`penilaian_id`)
        REFERENCES `penilaian` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_nilai_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE
);