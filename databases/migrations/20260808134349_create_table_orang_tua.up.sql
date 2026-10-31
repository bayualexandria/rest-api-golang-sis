CREATE TABLE `orang_tua` (
    id BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
   id_siswa INT UNIQUE NOT NULL,
   nik_ayah VARCHAR(100),
    nama_ayah VARCHAR(100),
    pekerjaan_ayah VARCHAR(50),
    nik_ibu VARCHAR(100),
    nama_ibu VARCHAR(100),
    pekerjaan_ibu VARCHAR(50),
    nama_wali VARCHAR(100),
    pekerjaan_wali VARCHAR(50),
    nomor_telepon VARCHAR(25) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);