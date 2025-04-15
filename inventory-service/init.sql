CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);
CREATE TABLE discounts (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    discount_percentage FLOAT NOT NULL,
    applicable_products JSONB,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    is_active BOOLEAN
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category_id INT NOT NULL,  -- заменяем category на category_id
    price NUMERIC(10, 2) NOT NULL,
    stock INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id)  -- внешняя ссылка на ID категории
);
