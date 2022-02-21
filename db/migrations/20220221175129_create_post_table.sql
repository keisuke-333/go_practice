-- +goose Up
CREATE TABLE posts (
  id int NOT NULL AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  body text,
  PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE posts;