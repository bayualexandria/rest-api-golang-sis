CREATE TABLE `absensi_siswa` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_kelas_id` BIGINT(20) ,
    `status_kehadiran_id` BIGINT(20) ,
    `tanggal` DATE NOT NULL,
    `keterangan` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_absensi_siswa_kelas`
        FOREIGN KEY (`siswa_kelas_id`)
        REFERENCES `siswa_kelas` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_absensi_status`
        FOREIGN KEY (`status_kehadiran_id`)
        REFERENCES `status_kehadiran` (`id`)
        ON DELETE RESTRICT
);