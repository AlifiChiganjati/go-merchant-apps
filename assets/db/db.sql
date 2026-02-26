CREATE DATABASE db_merchant;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  fullname VARCHAR(225) NOT NULL,
  password VARCHAR(225) NOT NULL,
  email VARCHAR(225) UNIQUE NOT NULL,
  phone_number VARCHAR(15) NOT NULL,
  role VARCHAR(255) NOT NULL,
  point INTEGER DEFAULT 0 NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "merchants" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID  REFERENCES users(id),
  merchant_name VARCHAR(30) NOT NULL,
  merchant_address VARCHAR(100) NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "products" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  merchant_id UUID  REFERENCES merchants(id),
  name VARCHAR(30) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  description VARCHAR(255) NOT NULL,
  point INTEGER NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "cart"(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID  REFERENCES users(id),
  product_id UUID REFERENCES products(id),
  quantity INTEGER NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transactions" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID  REFERENCES users(id),
  merchant_id UUID  REFERENCES merchants(id),
  transactions_no INTEGER NOT NULL,
  point_total INTEGER NOT NULL,
  subtotal DECIMAL(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS "transactions_detail" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  transaction_id UUID  REFERENCES transactions(id),
  product_id UUID  REFERENCES products(id),
  qty INTEGER NOT NULL,
  point INTEGER NOT NULL,
  total DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP  DEFAULT CURRENT_TIMESTAMP
);
