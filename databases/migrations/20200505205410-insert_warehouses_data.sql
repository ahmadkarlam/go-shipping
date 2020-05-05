-- +migrate Up
INSERT INTO go_shipping.warehouses
    (id, code, stock, x, y, created_at, updated_at)
VALUES (1, 'A', 10, 5, 18, '2020-05-05 13:14:03', null),
       (2, 'B', 10, 19, 17, '2020-05-05 13:14:03', null),
       (3, 'C', 10, 10, 18, '2020-05-05 13:14:03', null),
       (4, 'D', 10, 21, 14, '2020-05-05 13:14:03', null),
       (5, 'E', 10, 8, 10, '2020-05-05 13:14:03', null),
       (6, 'F', 10, 16, 9, '2020-05-05 13:14:03', null),
       (7, 'G', 10, 11, 5, '2020-05-05 13:14:03', null),
       (8, 'H', 10, 17, 4, '2020-05-05 13:14:03', null);

-- +migrate Down
TRUNCATE TABLE `warehouses`;
