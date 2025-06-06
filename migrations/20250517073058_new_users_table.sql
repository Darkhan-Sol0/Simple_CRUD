-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) UNIQUE NOT NULL,
    email VARCHAR(100)
);

INSERT INTO users (name, email) VALUES 
  ('firstTestUser', 'firstUser@email.com'),
  ('secondTestUser', 'secondUser@email.com'),
  ('thirdTestUser', 'thirdUser@email.com');
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd