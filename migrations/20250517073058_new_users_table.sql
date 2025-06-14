-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) UNIQUE NOT NULL,
    role VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(100)
);

INSERT INTO users (name, role, email, password) VALUES 
  ('da', 'admin', 'ddd@ddd.com', '$2a$10$SRKg/gRJKYjD3x/WaswPROtC072eLJkVYeIU5fiJmrzNHXALfPlZO'); 
-- password: qwe123qwe123
  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd