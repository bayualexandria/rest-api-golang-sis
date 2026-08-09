CREATE TABLE `orang_tua` (
    `id` BIGINT(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    `nik` VARCHAR(30),
    `nama` VARCHAR(255) NOT NULL,
    `jenis_kelamin` VARCHAR(10),
    `tempat_lahir` VARCHAR(100),
    `tanggal_lahir` DATE,
    `pendidikan` VARCHAR(100),
    `pekerjaan` VARCHAR(100),
    `no_hp` VARCHAR(30),
    `email` VARCHAR(255),
    `alamat` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
);

       