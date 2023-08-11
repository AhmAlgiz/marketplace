CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    pass_hash VARCHAR(255) NOT NULL
);

CREATE TABLE items
(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price MONEY NOT NULL,
    user_id INTEGER REFERENCES users (id) ON DELETE CASCADE
);
