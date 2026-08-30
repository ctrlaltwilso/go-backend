CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(500) NOT NULL,
  created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, description) VALUES
  ('Mechanical Keyboard', 'The keyboard is mechanical and goes click click.'),
  ('Mouse', 'Its a mouse but its not alive.')
