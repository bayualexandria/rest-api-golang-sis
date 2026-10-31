CREATE TABLE calon_siswa (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    id_user_calon_siswa BIGINT,
    nisn VARCHAR(10) UNIQUE NOT NULL,
    nik VARCHAR(16) UNIQUE NOT NULL,
    nama_lengkap VARCHAR(100) NOT NULL,
    ttl VARCHAR(255) NOT NULL,
    jenis_kelamin ENUM('Laki-laki', 'Perempuan') NOT NULL,
    agama VARCHAR(20) NOT NULL,
    alamat TEXT NOT NULL,
    sekolah_asal VARCHAR(100) NOT NULL,
    tahun_lulus YEAR NOT NULL,
     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  
    CONSTRAINT `fk_calon_siswa_user`
        FOREIGN KEY (`id_user_calon_siswa`)
        REFERENCES `user_calon_siswa` (`id`)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);