CREATE TABLE `siswa_orang_tua` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `orang_tua_id` BIGINT(20),
    `hubungan` VARCHAR(30) NOT NULL,
    `is_primary` TINYINT(1) NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_siswa_orang_tua_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_siswa_orang_tua_orang_tua`
        FOREIGN KEY (`orang_tua_id`)
        REFERENCES `orang_tua` (`id`)
        ON DELETE CASCADE
);