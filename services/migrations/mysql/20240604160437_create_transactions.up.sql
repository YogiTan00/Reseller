CREATE TABLE transactions (
    id                    VARCHAR(225) NOT NULL UNIQUE,
    id_name               VARCHAR(255),
    transaction_number    VARCHAR(255),
    sales_date            VARCHAR(255),
    unit                  INT,
    discount              FLOAT,
    admin_fee             FLOAT,
    add_admin_fee         FLOAT,
    created_at            TIMESTAMP,
    updated_at            TIMESTAMP,
    PRIMARY KEY(id)
)