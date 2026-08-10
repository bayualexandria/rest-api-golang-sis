CREATE TABLE `status_siswa` (
    -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    `nama_status` VARCHAR(100) NOT NULL,
    `keterangan` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
);
