CREATE TABLE `siswa_kelas` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `kelas_id` BIGINT(20) ,
    `tahun_ajaran_id` BIGINT(20) ,
    `semester_id` BIGINT(20) ,
    `status` VARCHAR(50) NOT NULL DEFAULT 'aktif',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,

    CONSTRAINT `fk_siswa_kelas_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL,

    CONSTRAINT `fk_siswa_kelas_kelas`
        FOREIGN KEY (`kelas_id`)
        REFERENCES `kelas` (`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT `fk_siswa_kelas_tahun`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT `fk_siswa_kelas_semester`
        FOREIGN KEY (`semester_id`)
        REFERENCES `semester` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);
        