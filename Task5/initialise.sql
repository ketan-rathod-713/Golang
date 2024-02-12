CREATE SCHEMA httpnet;

-- Create table user inside httpnet schema
CREATE TABLE httpnet.user(
  id SERIAL,
  firstname VARCHAR(80) NOT NULL,
  lastname VARCHAR(80),
  email VARCHAR(80) UNIQUE NOT NULL,
  phone VARCHAR(20) UNIQUE,
  dateofbirth date,
  PRIMARY KEY (id)
);

-- Insert data inside user table
INSERT INTO httpnet.user(firstname, lastname, dateofbirth, email, phone) VALUES('Ketan', 'Rathod', '07-01-2003', 'ketanrtd1@gmail.com', '9725488060');

-- Get all users
SELECT *  FROM httpnet.user;
