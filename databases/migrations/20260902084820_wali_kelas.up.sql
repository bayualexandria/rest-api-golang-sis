CREATE TABLE wali_kelas (
    id BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,

    guru_wali_id BIGINT(20) ,
    kelas_id BIGINT(20) ,
    tahun_ajaran_id BIGINT(20) ,
    semester_id BIGINT(20) ,

    status ENUM('aktif', 'nonaktif') NOT NULL DEFAULT 'aktif',

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

  
    CONSTRAINT `fk_wali_kelas_guru_wali`
        FOREIGN KEY (`guru_wali_id`)
        REFERENCES `guru` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL,

    CONSTRAINT `fk_wali_kelas_kelas`
        FOREIGN KEY (`kelas_id`)
        REFERENCES `kelas` (`id`)
         ON UPDATE CASCADE
        ON DELETE SET NULL,

    CONSTRAINT `fk_wali_kelas_tahun_ajaran`
        FOREIGN KEY (`tahun_ajaran_id`)
        REFERENCES `tahun_ajaran` (`id`)
         ON UPDATE CASCADE
        ON DELETE SET NULL,
    CONSTRAINT `fk_wali_kelas_semester`
        FOREIGN KEY (`semester_id`)
        REFERENCES `semester` (`id`)
         ON UPDATE CASCADE
        ON DELETE SET NULL
);