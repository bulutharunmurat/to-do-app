DROP TABLE IF EXISTS to_do_table;
CREATE TABLE to_do_table
(
    id        INT AUTO_INCREMENT NOT NULL,
    task      VARCHAR(128) NOT NULL,
    is_done   VARCHAR(128) NOT NULL,
    task_date VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO to_do_table
    (task, is_done, task_date)
VALUES ('task1', 'NO', '2022-01-01'),
       ('task2', 'YES', '2022-01-01');