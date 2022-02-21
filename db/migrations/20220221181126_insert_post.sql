-- +goose Up
INSERT INTO posts (
  title,
  body
)
VALUES
('Ruby','プログラミング入門者にはちょうどいい書きやすさ'),
('PHP','書きやすくwordpressでも使われている'),
('Go','低レイヤーも扱い易いコンパイラ言語');

-- +goose Down
DELETE FROM posts WHERE id IN (1, 2, 3);
