ALTER TABLE guru ADD COLUMN status_guru_id BIGINT AFTER alamat;

ALTER TABLE guru ADD CONSTRAINT `fk_guru_status_guru_id`
        FOREIGN KEY (`status_guru_id`)
        REFERENCES `status_guru` (`id`)
        ON DELETE CASCADE;