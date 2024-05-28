DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
        id SERIAL PRIMARY KEY,
        price NUMERIC NOT NULL,
        tax NUMERIC NOT NULL
);