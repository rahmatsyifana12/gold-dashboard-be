CREATE TABLE IF NOT EXISTS users (
   id serial PRIMARY KEY,
   email VARCHAR(255) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   name VARCHAR(255),
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP,
   deleted_at TIMESTAMP
);