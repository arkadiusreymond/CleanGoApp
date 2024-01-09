-- 001_create_birds_table.up.sql
CREATE TABLE birds (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       color VARCHAR(255) NOT NULL
);
