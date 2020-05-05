-- +migrate Up
CREATE TABLE IF NOT EXISTS `warehouses`
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    code       VARCHAR(10) NOT NULL,
    stock      INT,
    x          INT,
    y          INT,
    created_at TIMESTAMP   NULL,
    updated_at TIMESTAMP   NULL
);

-- +migrate Down
DROP TABLE `warehouses`;
