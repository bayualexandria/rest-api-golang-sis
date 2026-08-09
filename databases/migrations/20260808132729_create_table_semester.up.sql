CREATE TABLE `semester` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `tahun_ajaran_id` BIGINT(20) NOT NULL,
    `nama_semester` VARCHAR(50) NOT NULL,
    `kode` VARCHAR(20) NOT NULL,
    `is_active` TINYINT(1) NOT NULL DEFAULT 0,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,

    CONSTRAINT `fk_semester_tahun_ajaran`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);