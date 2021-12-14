CREATE DATABASE dm;
USE dm;
CREATE TABLE global_tx (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    gid VARCHAR(64) NOT NULL,
    expire_at TIMESTAMP NOT NULL
);
CREATE TABLE local_tx (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    global_tx_id BIGINT NOT NULL,
    callback_url VARCHAR(512) NOT NULL,
    FOREIGN KEY (global_tx_id) REFERENCES global_tx(id)
);

CREATE DATABASE bank1;
USE bank1;
CREATE TABLE wallet (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    balance DECIMAL NOT NULL
);

CREATE DATABASE bank2;
USE bank2;
CREATE TABLE wallet (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    balance DECIMAL NOT NULL
);
