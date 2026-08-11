-- Menamabahkan field baru pada tabel absensi siswa
ALTER TABLE absensi_siswa ADD COLUMN semester_id BIGINT AFTER keterangan;
ALTER TABLE absensi_siswa ADD COLUMN jam_masuk TIME NULL AFTER keterangan;
ALTER TABLE absensi_siswa ADD COLUMN jam_keluar TIME NULL AFTER keterangan;
ALTER TABLE absensi_siswa ADD COLUMN photo VARCHAR(255) AFTER keterangan;

ALTER TABLE absensi_siswa ADD CONSTRAINT `fk_absensi_siswa_semester`
        FOREIGN KEY (`semester_id`)
        REFERENCES `semester` (`id`)
        ON DELETE CASCADE;