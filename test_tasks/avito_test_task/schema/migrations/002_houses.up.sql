CREATE TABLE IF NOT EXISTS houses (
    id SERIAL PRIMARY KEY,
    address TEXT NOT NULL UNIQUE,
    year INT NOT NULL,
    developer TEXT DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);