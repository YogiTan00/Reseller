CREATE TABLE transactions (
    id                    VARCHAR(225) NOT NULL UNIQUE,
    id_name               VARCHAR(255),
    return_trans                BOOLEAN,
    sales_date            DATETIME DEFAULT NULL,
    unit                  INT,
    description               VARCHAR(255),
    created_at            DATETIME DEFAULT NULL,
    updated_at            DATETIME DEFAULT NULL,
    deleted_at            DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
)