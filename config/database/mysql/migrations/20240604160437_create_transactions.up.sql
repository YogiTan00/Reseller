CREATE TABLE transactions (
    id                    VARCHAR(225) NOT NULL UNIQUE,
    id_name               VARCHAR(255),
    return                BOOLEAN,
    sales_date            VARCHAR(255),
    unit                  INT,
    description               VARCHAR(255),
    created_at            DATETIME DEFAULT NULL,
    updated_at            DATETIME DEFAULT NULL,
    deleted_at            DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
)