CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nik VARCHAR(50) UNIQUE NOT NULL,
    full_name VARCHAR(100),
    legal_name VARCHAR(100),
    birth_place VARCHAR(50),
    birth_date DATE,
    salary DOUBLE PRECISION,
    photo_id_card TEXT,
    photo_selfie TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE limits (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    tenor_months INT,
    limit_amount DOUBLE PRECISION,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL,
    contract_number VARCHAR(100) UNIQUE NOT NULL,
    otr DOUBLE PRECISION,
    admin_fee DOUBLE PRECISION,
    installment_amount DOUBLE PRECISION,
    interest_amount DOUBLE PRECISION,
    asset_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP
);