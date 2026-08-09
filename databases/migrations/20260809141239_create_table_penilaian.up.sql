CREATE TABLE `penilaian` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `kelas_id` BIGINT(20) ,
    `mata_pelajaran_id` BIGINT(20) ,
    `guru_id` BIGINT(20) ,
    `jenis_penilaian_id` BIGINT(20) ,
    `semester_id` BIGINT(20) ,
    `nama_penilaian` VARCHAR(255) NOT NULL,
    `tanggal` DATE NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT `fk_penilaian_kelas`
        FOREIGN KEY (`kelas_id`) REFERENCES `kelas` (`id`),

    CONSTRAINT `fk_penilaian_mapel`
        FOREIGN KEY (`mata_pelajaran_id`) REFERENCES `mata_pelajaran` (`id`),

    CONSTRAINT `fk_penilaian_guru`
        FOREIGN KEY (`guru_id`) REFERENCES `guru` (`id`),

    CONSTRAINT `fk_penilaian_jenis`
        FOREIGN KEY (`jenis_penilaian_id`) REFERENCES `jenis_penilaian` (`id`),

    CONSTRAINT `fk_penilaian_semester`
        FOREIGN KEY (`semester_id`) REFERENCES `semester` (`id`)
);