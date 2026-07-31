CREATE TABLE master_saints (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    remark TEXT,
    status VARCHAR(50) DEFAULT 'active',
    created_at TIMESTAMPTZ DEFAULT now(),
    created_username VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT now(),
    updated_username VARCHAR(255)
);