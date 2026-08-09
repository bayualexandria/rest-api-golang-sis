CREATE TABLE `catatan_perkembangan` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `guru_id` BIGINT(20) ,
    `semester_id` BIGINT(20) ,
    `tanggal` DATE NOT NULL,
    `kategori` VARCHAR(100),
    `catatan` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_catatan_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_catatan_guru`
        FOREIGN KEY (`guru_id`)
        REFERENCES `guru` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_catatan_semester`
        FOREIGN KEY (`semester_id`)
        REFERENCES `semester` (`id`)
        ON DELETE CASCADE
);