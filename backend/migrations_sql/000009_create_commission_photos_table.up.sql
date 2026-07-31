CREATE TABLE commission_photos (
    id SERIAL PRIMARY KEY,
    commission_id INT NOT NULL REFERENCES commissions(id) ON DELETE CASCADE,
    file_url TEXT NOT NULL
);