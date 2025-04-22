CREATE TABLE "user" (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    role VARCHAR(20) DEFAULT 'user'
)