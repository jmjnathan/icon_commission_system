CREATE TABLE commission_items (
    id SERIAL PRIMARY KEY,
    commission_id INT NOT NULL REFERENCES commissions(id) ON DELETE CASCADE,
    saint_id INT NOT NULL REFERENCES master_saints(id),
    size_id INT NOT NULL REFERENCES master_sizes(id),
    material_id INT NOT NULL REFERENCES master_materials(id),
    style_id INT NOT NULL REFERENCES master_styles(id),
    price NUMERIC(15,2) DEFAULT 0,
    notes TEXT
);