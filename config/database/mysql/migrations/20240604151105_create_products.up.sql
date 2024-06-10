CREATE TABLE products (
    id              VARCHAR(225) NOT NULL UNIQUE,
    name            VARCHAR(255),
    type_size       VARCHAR(255),
    market_type     VARCHAR(255),
    image           VARCHAR(255),
    default_price   INT,
    price           INT,
    created_at      DATETIME,
    updated_at      DATETIME,
    PRIMARY KEY(id)
)