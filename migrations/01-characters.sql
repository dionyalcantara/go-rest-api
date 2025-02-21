CREATE DATABASE IF NOT EXISTS OverwatchAPI;

USE OverwatchAPI;

CREATE TABLE IF NOT EXISTS characters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon_url VARCHAR(255) NOT NULL,
    class VARCHAR(50) NOT NULL,
    description TEXT
) ENGINE=INNODB;
