CREATE TABLE commissions (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES clients(id),
    order_date TIMESTAMPTZ,
    deadline TIMESTAMPTZ,
    notes TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    total_price NUMERIC(15,2) DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(),
    created_username VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT now(),
    updated_username VARCHAR(255)
);