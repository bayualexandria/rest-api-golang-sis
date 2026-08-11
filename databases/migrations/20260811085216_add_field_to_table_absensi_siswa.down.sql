-- Drop fields from the table absensi siswa
ALTER TABLE absensi_siswa DROP FOREIGN KEY fk_absensi_siswa_semester; 
ALTER TABLE absensi_siswa DROP COLUMN IF EXISTS semester_id;
ALTER TABLE absensi_siswa DROP COLUMN IF EXISTS jam_masuk;
ALTER TABLE absensi_siswa DROP COLUMN IF EXISTS jam_keluar;
ALTER TABLE absensi_siswa DROP COLUMN IF EXISTS photo;
