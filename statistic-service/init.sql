CREATE TABLE IF NOT EXISTS order_statistics (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    items_count INTEGER NOT NULL,
    total_amount DECIMAL(10,2) NOT NULL,
    event_time TIMESTAMP NOT NULL,
    hour_of_day INTEGER NOT NULL,
    day_of_week INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS user_statistics (
    user_id VARCHAR(255) PRIMARY KEY,
    total_orders INTEGER DEFAULT 0,
    total_items_ordered INTEGER DEFAULT 0,
    favorite_category VARCHAR(255),
    average_order_value DECIMAL(10,2) DEFAULT 0,
    last_order_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS inventory_statistics (
    product_id VARCHAR(255) PRIMARY KEY,
    category_id VARCHAR(255) NOT NULL,
    stock_changes INTEGER DEFAULT 0,
    last_updated TIMESTAMP,
    times_out_of_stock INTEGER DEFAULT 0
);