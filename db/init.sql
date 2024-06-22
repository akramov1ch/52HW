CREATE TABLE sales_transactions (
    transaction_id VARCHAR PRIMARY KEY,
    product_id VARCHAR,
    quantity INT,
    price DECIMAL,
    timestamp BIGINT
);
