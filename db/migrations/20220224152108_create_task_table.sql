-- +goose Up
CREATE TABLE tasks
(
  id   int(11) UNSIGNED AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS tasks;