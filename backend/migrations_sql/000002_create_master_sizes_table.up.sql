CREATE TABLE master_sizes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'Active',
    created_at TIMESTAMPTZ DEFAULT now(),
    created_username VARCHAR(255),
    updated_at TIMESTAMPTZ DEFAULT now(),
    updated_username VARCHAR(255)
);