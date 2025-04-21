CREATE TABLE products (
    id              VARCHAR(225) NOT NULL UNIQUE,
    name            VARCHAR(255),
    type_size       VARCHAR(255),
    image           VARCHAR(255),
    default_price   INT DEFAULT 0,
    price           INT DEFAULT 0,
    created_at      DATETIME DEFAULT NOW(),
    updated_at      DATETIME DEFAULT NOW(),
    deleted_at      DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
)