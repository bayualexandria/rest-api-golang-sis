ALTER TABLE siswa_kelas ADD COLUMN wali_kelas_id BIGINT AFTER semester_id;

ALTER TABLE siswa_kelas ADD CONSTRAINT `fk_siswa_kelas_wali_kelas`
        FOREIGN KEY (`wali_kelas_id`)
        REFERENCES `guru` (`id`)
        ON DELETE CASCADE;