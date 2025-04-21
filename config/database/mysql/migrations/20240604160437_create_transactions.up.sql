CREATE TABLE transactions (
    id                    VARCHAR(225) NOT NULL UNIQUE,
    id_name               VARCHAR(255),
    return_trans          BOOLEAN DEFAULT FALSE,
    sales_date            DATETIME DEFAULT NOW(),
    unit                  INT DEFAULT 0,
    description           VARCHAR(255),
    created_at            DATETIME DEFAULT NOW(),
    updated_at            DATETIME DEFAULT NOW(),
    deleted_at            DATETIME DEFAULT NULL,
    PRIMARY KEY(id)
)