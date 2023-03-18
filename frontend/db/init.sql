CREATE DATABASE postgres;
\c postgres;
CREATE TABLE tb_casestudy (
  ID SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  imageuri VARCHAR(255) NOT NULL,
  createddate TIMESTAMP NOT NULL DEFAULT NOW()
);

--  postgres adli veritabanı oluşturur ve tb_casestudy adli table oluşturur
