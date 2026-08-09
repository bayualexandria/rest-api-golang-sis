CREATE TABLE `guru_mata_pelajaran` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `guru_id` BIGINT(20) ,
    `mata_pelajaran_id` BIGINT(20) ,
    `tahun_ajaran_id` BIGINT(20) ,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_gmp_guru`
        FOREIGN KEY (`guru_id`)
        REFERENCES `guru` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_gmp_mapel`
        FOREIGN KEY (`mata_pelajaran_id`)
        REFERENCES `mata_pelajaran` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_gmp_tahun`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
        ON DELETE CASCADE
);