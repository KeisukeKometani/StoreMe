BEGIN;

  CREATE TABLE IF NOT EXISTS transaction_products (
    -- COMMON SQL
    -- id          INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id          INTEGER NOT NULL AUTO_INCREMENT PRIMARY KEY,
    product_id       INTEGER NOT NULL,
    transaction_at TIMESTAMP NOT NULL,
    record_at TIMESTAMP NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP NULL DEFAULT NULL,

    FOREIGN KEY fk_transaction_products_product_id (`product_id`)
      REFERENCES products (`id`)
  );

COMMIT;
