CREATE TABLE cash_outs (
    id SERIAL PRIMARY KEY,
    category VARCHAR(255),
    description TEXT,
    qty NUMERIC(15,2) DEFAULT 0,
    unit_price NUMERIC(15,2) DEFAULT 0,
    amount NUMERIC(15,2) DEFAULT 0,
    vendor VARCHAR(255),
    payment_method VARCHAR(100),
    receipt_url TEXT,
    notes TEXT,
    date TIMESTAMPTZ,
    status VARCHAR(50) DEFAULT 'Active',
    created_at TIMESTAMPTZ DEFAULT now(),
    created_username VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT now(),
    updated_username VARCHAR(255)
);