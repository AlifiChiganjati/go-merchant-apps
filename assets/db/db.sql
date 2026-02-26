CREATE DATABASE IF NOT EXISTS db_sms_ppob;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "users"(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(255) NOT NULL, 
  last_name VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  profile_image VARCHAR(255),
  saldo DECIMAL(10,2) DEFAULT 0,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "banners"(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  banner_name VARCHAR(255),
  banner_image VARCHAR(255),
  description VARCHAR(255),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "services"(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  service_code VARCHAR(255),
  service_name VARCHAR(255),
  service_icon VARCHAR(255),
  service_tarif DECIMAL(10,2),
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transactions"(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id),
  service_id UUID REFERENCES services(id),
  transaction_type VARCHAR(255) NOT NULL,
  invoice_number VARCHAR(255) NOT NULL,
  total_amount DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

INSERT INTO "banners" 
  (banner_name,banner_image,description,created_at,updated_at)
VALUES
  ('Banner 1','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp),
  ('Banner 2','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp),
  ('Banner 3','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp),
  ('Banner 4','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp),
  ('Banner 5','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp),
  ('Banner 6','https://placehold.co/600x400','Lorem Ipsum Dolor sit amet',current_timestamp,current_timestamp);

INSERT INTO "services"
  (service_code, service_name, service_icon, service_tarif,created_at,updated_at)
VALUES 
  ('PAJAK','Pajak PBB','https://placehold.co/600x400',40000,current_timestamp,current_timestamp),
  ('PLN','Listrik','https://placehold.co/600x400',10000,current_timestamp,current_timestamp),
  ('PDAM','PDAM Berlangganan','https://placehold.co/600x400',40000,current_timestamp,current_timestamp),
  ('PULSA','Pulsa','https://placehold.co/600x400',40000,current_timestamp,current_timestamp),
  ('PGN','PGN Berlangganan','https://placehold.co/600x400',50000,current_timestamp,current_timestamp),
  ('MUSIK','Musik Berlangganan','https://placehold.co/600x400',50000,current_timestamp,current_timestamp),
  ('TV','TV Berlangganan','https://placehold.co/600x400',50000,current_timestamp,current_timestamp),
  ('PAKET_DATA','Paket Data','https://placehold.co/600x400',50000,current_timestamp,current_timestamp),
  ('VOUCHER_GAME','Voucher Game','https://placehold.co/600x400',100000,current_timestamp,current_timestamp),
  ('VOUCHER_MAKANAN','Voucher Makanan','https://placehold.co/600x400',100000,current_timestamp,current_timestamp),
  ('QURBAN','Qurban','https://placehold.co/600x400',200000,current_timestamp,current_timestamp),
  ('ZAKAT','Zakat','https://placehold.co/600x400',300000,current_timestamp,current_timestamp);
