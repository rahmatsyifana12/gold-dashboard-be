CREATE TABLE gold_assets (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    brand gold_brand NOT NULL,
    unit_gram NUMERIC(10,4) NOT NULL CHECK (unit_gram > 0),
    certificate_no VARCHAR(100),
    status gold_status NOT NULL DEFAULT 'owned',
    bought_price NUMERIC(18,2) NOT NULL CHECK (bought_price >= 0),
    sold_price NUMERIC(18,2),
    buy_date DATE NOT NULL,
    sell_date DATE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);
