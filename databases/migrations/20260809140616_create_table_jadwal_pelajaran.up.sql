CREATE TABLE `jadwal_pelajaran` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `kelas_id` BIGINT(20) ,
    `guru_id` BIGINT(20) ,
    `mata_pelajaran_id` BIGINT(20) ,
    `tahun_ajaran_id` BIGINT(20) ,
    `hari` VARCHAR(20) NOT NULL,
    `jam_mulai` TIME NOT NULL,
    `jam_selesai` TIME NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL,

    CONSTRAINT `fk_jadwal_kelas`
        FOREIGN KEY (`kelas_id`)
        REFERENCES `kelas` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_jadwal_guru`
        FOREIGN KEY (`guru_id`)
        REFERENCES `guru` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_jadwal_mapel`
        FOREIGN KEY (`mata_pelajaran_id`)
        REFERENCES `mata_pelajaran` (`id`)
        ON DELETE CASCADE,

    CONSTRAINT `fk_jadwal_tahun`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
        ON DELETE CASCADE
);