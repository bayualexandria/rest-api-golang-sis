CREATE TABLE `pembayaran_siswa` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `siswa_id` BIGINT(20) ,
    `jenis_pembayaran_id` BIGINT(20) ,
    `tahun_ajaran_id` BIGINT(20) ,
    `tanggal_pembayaran` DATE,
    `nominal` DECIMAL(15,2) NOT NULL,
    `status` VARCHAR(30) NOT NULL DEFAULT 'belum_lunas',
    `keterangan` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_pembayaran_siswa`
        FOREIGN KEY (`siswa_id`)
        REFERENCES `siswa` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_pembayaran_jenis`
        FOREIGN KEY (`jenis_pembayaran_id`)
        REFERENCES `jenis_pembayaran` (`id`),

    CONSTRAINT `fk_pembayaran_tahun`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
);