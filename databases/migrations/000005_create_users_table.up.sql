CREATE TABLE users (
    -- Jika menggunakan MySQL, gunakan 'id INT AUTO_INCREMENT PRIMARY KEY' untuk kolom id
    -- Jika menggunakan PostgreSQL, gunakan 'id SERIAL PRIMARY KEY' atau BIGSERIAL untuk kolom id
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    username BIGINT NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE ,
    email_verified_at TIMESTAMP NULL,
    password VARCHAR(255) NOT NULL,
    status_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

        CONSTRAINT `fk_user_status_id`
        FOREIGN KEY (`status_id`)
        REFERENCES `status_user` (`id`)
        ON DELETE CASCADE
);