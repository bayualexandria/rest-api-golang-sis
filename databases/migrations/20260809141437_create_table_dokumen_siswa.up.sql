CREATE TABLE `dokumen_siswa` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `jenis_dokumen` VARCHAR(100) NOT NULL,
    `nama_file` VARCHAR(255) NOT NULL,
    `file_path` VARCHAR(500) NOT NULL,
    `keterangan` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_dokumen_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE
);