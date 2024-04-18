CREATE TABLE accounts(
    id UUID PRIMARY KEY,
    phone VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(1024) NOT NULL,
    name VARCHAR(64),
    profile_picture VARCHAR(2048)
);