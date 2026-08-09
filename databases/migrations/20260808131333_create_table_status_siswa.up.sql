CREATE TABLE `status_siswa` (
    -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nama_status` VARCHAR(100) NOT NULL,
    `keterangan` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
);
ALTER TABLE `siswa`
ADD CONSTRAINT `fk_siswa_status`
FOREIGN KEY (`status_siswa_id`)
REFERENCES `status_siswa` (`id`)
ON UPDATE CASCADE
ON DELETE SET NULL;