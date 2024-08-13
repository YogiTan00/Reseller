CREATE TABLE products (
    id              VARCHAR(225) NOT NULL UNIQUE,
    name            VARCHAR(255),
    type_size       VARCHAR(255),
    image           VARCHAR(255),
    default_price   INT,
    price           INT,
    created_at      DATETIME DEFAULT NULL,
    updated_at      DATETIME DEFAULT NULL,
    deleted_at      DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
)