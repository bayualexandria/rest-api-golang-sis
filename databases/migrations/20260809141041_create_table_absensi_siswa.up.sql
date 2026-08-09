CREATE TABLE `absensi_siswa` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `kelas_id` BIGINT(20) ,
    `status_kehadiran_id` BIGINT(20) ,
    `tanggal` DATE NOT NULL,
    `keterangan` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_absensi_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_absensi_kelas`
        FOREIGN KEY (`kelas_id`)
        REFERENCES `kelas` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_absensi_status`
        FOREIGN KEY (`status_kehadiran_id`)
        REFERENCES `status_kehadiran` (`id`)
        ON DELETE RESTRICT
);