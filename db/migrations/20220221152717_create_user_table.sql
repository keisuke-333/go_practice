-- +goose Up
CREATE TABLE users
(
  id   int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name varchar(255) NOT NULL
);

-- +goose Down
DROP TABLE users;