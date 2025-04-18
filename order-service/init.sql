CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,            -- ID заказа (генерируется автоматически)
    user_id VARCHAR(255) NOT NULL,    -- ID пользователя
    items TEXT[] NOT NULL,            -- Список товаров (массив строк, т.к. каждый товар представлен ID)
    total DECIMAL(10, 2) NOT NULL,    -- Общая сумма заказа
    status VARCHAR(50) NOT NULL,      -- Статус заказа
    created_at TIMESTAMP NOT NULL     -- Время создания заказа
);

-- Индексы для ускорения поиска
CREATE INDEX idx_user_id ON orders(user_id);
CREATE INDEX idx_status ON orders(status);
